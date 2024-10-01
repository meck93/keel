# Deployment

1. Create two secrets in the cluster:

```bash
kubectl create secret docker-registry ghcr-cred --docker-server=https://ghcr.io --docker-username=<username> --docker-password=<GITHUB_PAT> --docker-email=<email> -n keel
kubectl create secret generic keel-slack-cred --from-literal=app_token='xapp-...' --from-literal=bot_token='xoxb-...' --from-literal=channels='general' --from-literal=approvals_channel='general' --from-literal=bot_name='keel'
kubectl create secret generic keel-basic-auth-cred --from-literal=username='' --from-literal=password=''
```

2. Deploy keel

```bash
kubectl apply -f deployment.yaml
```
