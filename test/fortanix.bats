#!/usr/bin/env bats
WAIT_TIME=120
SLEEP_TIME=1

setup() { 
  if [[ -z "${FORTANIX_API_KEY}" ]]; then 
    echo "Error: Please provide fortanix api key" >&2 
    return 1 
  fi 
} 
@test "install fortanix csi provider" {
  # install the fortanix csi provider with the deployment file
  kubectl apply -f ../deployment/fortanix-csi-provider.yaml --namespace kube-system
  # wait for akeyless and akeyless-csi-provider pods to be running
  kubectl wait --for=condition=Ready --timeout=${WAIT_TIME}s pods --all -n kube-system
}
