#!/usr/bin/env bats

load 'bats-support/load'
load 'bats-assert/load'

@test "fortanix provider" {
  echo "Starting Fortanix CSI provider test..." >&3

  # Skip if required environment variables are not set
  [[ -n "${FORTANIX_API_KEY}" ]] || skip "FORTANIX_API_KEY not set"
  [[ -n "${FORTANIX_DSM_ENDPOINT}" ]] || skip "FORTANIX_DSM_ENDPOINT not set"
  [[ -n "${FORTANIX_TEST_SECRET_NAME}" ]] || skip "FORTANIX_TEST_SECRET_NAME not set"

  echo "Creating API key secret..." >&3
  cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: fortanix-api-key
  namespace: kube-system
type: Opaque
stringData:
  api-key: ${FORTANIX_API_KEY}
EOF

  echo "Deploying Fortanix CSI provider..." >&3
  kubectl apply -f deployment/fortanix-csi-provider.yaml

  echo "Waiting for provider pods to be ready..." >&3
  kubectl wait --for=condition=ready pod -l app=fortanix-csi-provider --namespace kube-system --timeout=300s

  echo "Creating SecretProviderClass..." >&3
  cat <<EOF | kubectl apply -f -
apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: fortanix-test
  namespace: kube-system
spec:
  provider: fortanix-csi-provider
  parameters:
    dsmEndpoint: "${FORTANIX_DSM_ENDPOINT}"
    objects: |
      - secretName: "${FORTANIX_TEST_SECRET_NAME}"
EOF

  echo "Creating test pod with CSI volume..." >&3
  cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: fortanix-test-pod
  namespace: kube-system
spec:
  containers:
  - name: test
    image: busybox
    command: ["sh", "-c", "sleep 3600"]
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
        secretProviderClass: fortanix-test
EOF

  echo "Waiting for test pod to be ready..." >&3
  kubectl wait --for=condition=ready pod fortanix-test-pod --namespace kube-system --timeout=300s

  echo "Checking if secret is mounted..." >&3
  run kubectl exec fortanix-test-pod --namespace kube-system -- ls /mnt/secrets
  assert_success
  assert_output --partial "${FORTANIX_TEST_SECRET_NAME}"
  echo "✓ Secret file found: ${FORTANIX_TEST_SECRET_NAME}" >&3

  echo "Checking secret content..." >&3
  run kubectl exec fortanix-test-pod --namespace kube-system -- cat "/mnt/secrets/${FORTANIX_TEST_SECRET_NAME}"
  assert_success
  [[ -n "$output" ]]
  echo "✓ Secret content retrieved successfully" >&3

  echo "Cleaning up resources..." >&3
  kubectl delete pod fortanix-test-pod --namespace kube-system --ignore-not-found=true
  kubectl delete secretproviderclass fortanix-test --namespace kube-system --ignore-not-found=true
  kubectl delete secret fortanix-api-key --namespace kube-system --ignore-not-found=true

  echo "Test completed successfully!" >&3
}
