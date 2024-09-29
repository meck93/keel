package policies

import (
	"github.com/meck93/keel/types"
)

// GetTriggerPolicy - checks for trigger label, if not set - returns
// default trigger type
func GetTriggerPolicy(labels map[string]string, annotations map[string]string) types.TriggerType {

	triggerAnn, ok := annotations[types.KeelTriggerLabel]
	if ok {
		return types.ParseTrigger(triggerAnn)
	}

	// checking labels
	trigger, ok := labels[types.KeelTriggerLabel]
	if ok {
		return types.ParseTrigger(trigger)
	}

	return types.TriggerTypeDefault
}
