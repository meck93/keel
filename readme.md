<p align="center">
  <a href="https://keel.sh" target="_blank"><img width="100"src="https://keel.sh/img/logo.png"></a>
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/keel-hq/keel">
    <img src="https://goreportcard.com/badge/github.com/keel-hq/keel" alt="Go Report">
  </a>
  
  <a href="https://img.shields.io/docker/pulls/keelhq/keel.svg">
    <img src="https://img.shields.io/docker/pulls/keelhq/keel.svg" alt="Docker Pulls">
  </a>

  <a href="https://drone-kr.webrelay.io/meck93/keel">
    <img src="https://drone-kr.webrelay.io/api/badges/meck93/keel/status.svg" alt="Drone Status">
  </a>
</p>

# Keel - automated Kubernetes deployments for the rest of us

Keel is a tool for automating [Kubernetes](https://kubernetes.io/) deployment updates. Keel is stateless, robust and lightweight.

Keel provides several key features:

- **[Kubernetes](https://kubernetes.io/) and [Helm](https://helm.sh) providers** - Keel has direct integrations with Kubernetes and Helm.

- **No CLI/API** - tired of `f***ctl` for everything? Keel doesn't have one. Gets job done through labels, annotations, charts.

- **Semver policies** - specify update policy for each deployment/Helm release individually.

- **Automatic [Google Container Registry](https://cloud.google.com/container-registry/) configuration** - Keel automatically sets up topic and subscriptions for your deployment images by periodically scanning your environment.

- **[Native, DockerHub, Quay and Azure container registry webhooks](https://keel.sh/docs/#triggers) support** - once webhook is received impacted deployments will be identified and updated.

- **[Polling](https://keel.sh/docs/#polling)** - when webhooks and pubsub aren't available - Keel can still be useful by checking Docker Registry for new tags (if current tag is semver) or same tag SHA digest change (ie: `latest`).

- **Notifications** - out of the box Keel has Slack and standard webhook notifications, more info [here](https://keel.sh/docs/#notifications)

<p align="center">
  <a href="https://keel.sh" target="_blank"><img width="700"src="https://keel.sh/img/keel_high_level.png"></a>
</p>

### Helm quick start

Prerequisites:

- Kubernetes

[Deployment](./deployment/README.md)

[Configuration](https://github.com/keel-hq/keel#configuration)

### Quick Start

<p align="center">
  <a href="https://keel.sh" target="_blank"><img width="700"src="https://keel.sh/img/examples/force-workflow.png"></a>
</p>

A step-by-step guide to install Keel on your Kubernetes cluster is viewable on the Keel website:

[https://keel.sh/examples/#example-1-push-to-deploy](https://keel.sh/examples/#example-1-push-to-deploy)

### Configuration

Once Keel is deployed, you only need to specify update policy on your deployment file or Helm chart:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wd
  namespace: default
  labels:
    name: "wd"
  annotations:
    keel.sh/policy: minor # <-- policy name according to https://semver.org/
    keel.sh/trigger: poll # <-- actively query registry, otherwise defaults to webhooks
spec:
  template:
    metadata:
      name: wd
      labels:
        app: wd
    spec:
      containers:
        - image: karolisr/webhook-demo:0.0.8
          imagePullPolicy: Always
          name: wd
          command: ["/bin/webhook-demo"]
          ports:
            - containerPort: 8090
```

No additional configuration is required. Enabling continuous delivery for your workloads has never been this easy!

### Documentation

Documentation is viewable [here](./documentation.md)

### Developing Keel

If you wish to work on Keel itself, you will need Go 1.23+ installed. Make sure you put Keel into correct Gopath and `go build` (dependency management is done through [dep](https://github.com/golang/dep)).

To test Keel while developing:

1. Launch a Kubernetes cluster like Minikube or Docker for Mac with Kubernetes.
2. Change config to use it: `kubectl config use-context docker-for-desktop`
3. Build Keel from `cmd/keel` directory.
4. Start Keel with: `keel --no-incluster`. This will use Kubeconfig from your home.

### Running unit tests

Get a test parser (makes output nice):

```bash
go get github.com/mfridman/tparse
```

To run unit tests:

```bash
make test
```

### Running e2e tests

Prerequisites:

- configured kubectl + kubeconfig
- a running cluster (test suite will create testing namespaces and delete them after tests)
- Go environment (will compile Keel before running)

Once prerequisites are ready:

```bash
make e2e
```
