package helm3

import (
	"github.com/meck93/keel/internal/policy"
	"github.com/meck93/keel/types"
	"github.com/meck93/keel/util/image"

	hapi_chart "helm.sh/helm/v3/pkg/chart"

	log "github.com/sirupsen/logrus"
)

func checkRelease(repo *types.Repository, namespace, name string, chart *hapi_chart.Chart, config map[string]interface{}) (plan *UpdatePlan, shouldUpdateRelease bool, err error) {

	plan = &UpdatePlan{
		Chart:       chart,
		Namespace:   namespace,
		Name:        name,
		Values:      make(map[string]string),
		EmptyConfig: config == nil,
	}

	eventRepoRef, err := image.Parse(repo.String())
	if err != nil {
		log.WithFields(log.Fields{
			"error":           err,
			"repository_name": repo.Name,
		}).Error("provider.helm3: failed to parse event repository name")
		return
	}

	// getting configuration
	vals, err := values(chart, config)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("provider.helm3: failed to get values.yaml for release")
		return
	}

	keelCfg, err := getKeelConfig(vals)
	if err != nil {
		if err == ErrPolicyNotSpecified {
			// nothing to do
			return plan, false, nil
		}
		log.WithFields(log.Fields{
			"error": err,
		}).Error("provider.helm3: failed to get keel configuration for release")
		// ignoring this release, no keel config found
		return plan, false, nil
	}
	log.Infof("policy for release %s/%s parsed: %s", namespace, name, keelCfg.Plc.Name())

	if keelCfg.Plc.Type() == policy.PolicyTypeNone {
		// policy is not set, ignoring release
		return plan, false, nil
	}

	// checking for impacted images
	for _, imageDetails := range keelCfg.Images {
		imageRef, err := parseImage(vals, &imageDetails)
		if err != nil {
			log.WithFields(log.Fields{
				"error":           err,
				"repository_name": imageDetails.RepositoryPath,
				"repository_tag":  imageDetails.TagPath,
			}).Error("provider.helm3: failed to parse image")
			continue
		}

		if imageRef.Repository() != eventRepoRef.Repository() {
			log.WithFields(log.Fields{
				"parsed_image_name": imageRef.Remote(),
				"target_image_name": repo.Name,
			}).Debug("provider.helm3: images do not match, ignoring")
			continue
		}

		shouldUpdate, err := keelCfg.Plc.ShouldUpdate(imageRef.Tag(), eventRepoRef.Tag())
		if err != nil {
			log.WithFields(log.Fields{
				"error":           err,
				"repository_name": imageDetails.RepositoryPath,
				"repository_tag":  imageDetails.TagPath,
			}).Error("provider.helm3: got error while checking whether update the chart")
			continue
		}

		if !shouldUpdate {
			log.WithFields(log.Fields{
				"parsed_image_name": imageRef.Remote(),
				"target_image_name": repo.Name,
				"policy":            keelCfg.Plc.Name(),
			}).Info("provider.helm3: ignoring")
			continue
		}

		if imageDetails.DigestPath != "" {
			plan.Values[imageDetails.DigestPath] = repo.Digest
			log.WithFields(log.Fields{
				"image_details_digestPath": imageDetails.DigestPath,
				"target_image_digest":      repo.Digest,
			}).Debug("provider.helm3: setting image Digest")
		}

		path, value := getUnversionedPlanValues(repo.Tag, imageRef, &imageDetails)
		plan.Values[path] = value
		plan.NewVersion = repo.Tag
		plan.CurrentVersion = imageRef.Tag()
		plan.Config = keelCfg
		shouldUpdateRelease = true
		if imageDetails.ReleaseNotes != "" {
			plan.ReleaseNotes = append(plan.ReleaseNotes, imageDetails.ReleaseNotes)
		}

	}

	return plan, shouldUpdateRelease, nil
}
