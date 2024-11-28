(this readme is for internal use only)

Fortanix provider for the
[Secrets Store CSI driver](https://github.com/kubernetes-sigs/secrets-store-csi-driver)
allows you to fetch/ secrets stored in Fortanix DSM and use the Secrets Store
CSI driver interface to mount them into Kubernetes pods.

# Installation

### Prerequisites

- Kubernetes 1.16+ for both the master and worker nodes (Linux-only)
- [Secrets store CSI driver](https://secrets-store-csi-driver.sigs.k8s.io/getting-started/installation.html)
  installed

### Using yaml

You can also install using the deployment config in the `deployment` folder:

```bash
kubectl apply -f deployment/fortanix-csi-provider.yaml
```

## To Enable Secret Rotation :

Use flag : `--set syncSecret.enabled=true \`--set
syncSecret.retryDuration=<duration>`

The `<duration>` can be in mins (2m) or in seconds (120s)

```
helm install csi-secrets-store secrets-store-csi-driver/secrets-store-csi-driver \
  --namespace kube-system \
  --set syncSecret.enabled=true \
  --set syncSecret.retryDuration=2m
```

## To Enable Sync as Kubernetes Secret :

Use :

```
helm install secrets-store-csi-driver secrets-store-csi-driver/secrets-store-csi-driver --namespace kube-system --set syncSecret.enabled=true
```

Add the following to your secret provider class yaml file :

```
secretObjects:
    - secretName: "<secret_name>"  #The Kubernetes secret name
      type: Opaque #type of kubernetes secret (tls,sat,basic etc..,)
      labels:
        environment: "test"
      data:
        - objectName: db-string  #This should match the secret name given in Fortanix Provider
          key: <key-value pair>
```

### To use as Environment Variable :

```
env:
      - name: <env_var_name>
        valueFrom:
          secretKeyRef:
            name: <name>           # Reference the created Kubernetes secret
            key: <key-value pair>  # Reference the key in the secret
```

## To Enable Both :

```
helm install secrets-store-csi-driver secrets-store-csi-driver/secrets-store-csi-driver --namespace kube-system  --set enableSecretRotation=true --set syncSecret.enabled=true --set syncSecret.retryDuration=1m
```

## To Upgrade An Existing Helm Installation:

```
helm upgrade secrets-store-csi-driver secrets-store-csi-driver/secrets-store-csi-driver  --namespace kube-system --set syncSecret.enabled=true --set syncSecret.enabled=true --set syncSecret.retryDuration=1m
```

## Troubleshooting :

To troubleshoot issues with the CSI provider, look at logs from the CSI provider
pod. steps for trobuleshooting :

1. Check Pod Status and Events

First, get details about the DaemonSet pods and look for any errors:

`kubectl get pods -n csi -l app=fortanix-csi-provider`\
If the pod is not in the `Running` state, describe the pod to get detailed
information:\
`kubectl describe pod -n csi <fortanix-csi-provider-pod-name>`\
Look for events and error messages in the output, such
as `Failed`, `CrashLoopBackOff`, or `Error`

1. Inspect Pod Logs
2. Check the logs of the Fortanix CSI Provider pod to identify any issues that
   occurred during startup:
3. `kubectl logs -n csi <fortanix-csi-provider-pod-name>`
