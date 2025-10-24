# Fortanix CSI Provider

The Fortanix provider for the [Secrets Store CSI driver](https://github.com/kubernetes-sigs/secrets-store-csi-driver) allows you to fetch secrets stored in Fortanix DSM and use the Secrets Store CSI driver interface to mount them into Kubernetes pods.

## Installation

### Prerequisites

- Kubernetes 1.16+ for both the master and worker nodes (Linux-only)
- [Secrets Store CSI driver](https://secrets-store-csi-driver.sigs.k8s.io/getting-started/installation.html) installed
- Access to a Fortanix DSM instance with API key

### Setup API Key

Before deploying the provider, create a Kubernetes secret with your Fortanix API key:

```bash
kubectl create secret generic fortanix-api-key --from-literal=api-key=YOUR_FORTANIX_API_KEY -n kube-system
```

Replace `YOUR_FORTANIX_API_KEY` with your actual API key.

### Deploy Provider

Install using the deployment config in the `deployment` folder:

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

## Usage

### SecretProviderClass Configuration

Create a `SecretProviderClass` to define which secrets to fetch from Fortanix DSM:

```yaml
apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: fortanix-secret-provider
  namespace: default
spec:
  provider: fortanix-csi-provider
  parameters:
    dsmEndpoint: "https://your-dsm-endpoint.smartkey.io"
    objects: |
      - secretName: "my-secret"
```

### Mounting Secrets in Pods

Mount the secrets in your pods using the CSI volume:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: my-container
    image: nginx
    volumeMounts:
    - name: secrets
      mountPath: /mnt/secrets
      readOnly: true
  volumes:
  - name: secrets
    csi:
      driver: secrets-store.csi.k8s.io
      readOnly: true
      volumeAttributes:
        secretProviderClass: fortanix-secret-provider
```

## Testing

This provider includes end-to-end tests using BATS. See [test/bats/README.md](test/bats/README.md) for details on running the tests.

To run the tests:

```bash
make test-e2e-fortanix
```

## Troubleshooting

To troubleshoot issues with the CSI provider, review the logs from the CSI provider pod. Follow these steps for effective troubleshooting:

- Check Pod Status and Events: Get details about the DaemonSet pods and look for any errors:
  ```bash
  kubectl get pods -n kube-system -l app=fortanix-csi-provider
  ```
  If the pod is not in the `Running` state, describe the pod to get detailed information:
  ```bash
  kubectl describe pod -n kube-system <fortanix-csi-provider-pod-name>
  ```
  Look for events and error messages in the output, such as `Failed`, `CrashLoopBackOff`, or `Error`.

- Inspect Pod Logs: Check the logs of the Fortanix CSI Provider pod to identify any issues that occurred during startup:
  ```bash
  kubectl logs -n kube-system <fortanix-csi-provider-pod-name>
  ```
