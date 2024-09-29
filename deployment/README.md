# Deployment

1. Create two secrets in the cluster:

```bash
kubectl create secret generic keel-slack-cred --from-literal=app_token='' --from-literal=bot_token='' --from-literal=approvals_channel='' --from-literal=approvals_channel='' --from-literal=bot_name='keel'
kubectl create secret generic keel-basic-auth-cred --from-literal=username='' --from-literal=password=''
```

2. Deploy keel

```bash
kubectl apply -f deployment.yaml
```
