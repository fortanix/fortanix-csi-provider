#  * Copyright (c) Fortanix, Inc.
#  * This Source Code Form is subject to the terms of the Mozilla Public
#  * License, v. 2.0. If a copy of the MPL was not distributed with this
#  * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

#!/usr/bin/env bats

# Open FD 3 to stderr for debug/info prints
exec 3>&2

# ------------------------
# Config
# ------------------------
WAIT_TIME=180
SLEEP_TIME=2
# Use raw.githubusercontent URL so kubectl can apply it directly
PROVIDER_YAML=${PROVIDER_YAML:-https://raw.githubusercontent.com/fortanix/fortanix-csi-provider/main/deployment/fortanix-csi-provider.yaml}
NAMESPACE=${NAMESPACE:-kube-system}
POD_NAME=${POD_NAME:-fortanix-test-pod}
BATS_TEST_DIR=${BATS_TEST_DIR:-test/bats}

# Required runtime envs (the tests will skip if these are missing)
# Per README the provider reads a Kubernetes secret named `fortanix-api-key` with key `api-key`.
# Provide your Fortanix API key here before running tests.
export FORTANIX_API_KEY=${FORTANIX_API_KEY:-}
export FORTANIX_DSM_ENDPOINT=${FORTANIX_DSM_ENDPOINT:-}
export FORTANIX_TEST_SECRET_NAME=${FORTANIX_TEST_SECRET_NAME:-}

# ------------------------
# Minimal local helper functions (small subset of bats-assert)
# ------------------------
assert_success() {
  local got_status="${status:-0}"
  if [ "$got_status" -ne 0 ]; then
    echo "Assertion failed: expected status 0, got ${got_status}" >&2
    return 1
  fi
}

assert_output_contains() {
  local expected="$1"
  if ! printf '%s' "${output:-}" | grep -Fq -- "$expected"; then
    echo "Assertion failed: expected output to contain: ${expected}" >&2
    echo "Actual output:" >&2
    printf '%s\n' "${output:-}" >&2
    return 1
  fi
}

assert_failure() {
  local got_status="${status:-0}"
  if [ "$got_status" -eq 0 ]; then
    echo "Assertion failed: expected non-zero status, got 0" >&2
    return 1
  fi
}

# ------------------------
# Test utility helpers
# ------------------------
wait_for_process() {
  local timeout="$1"; shift
  local sleep_time="$1"; shift
  local cmd="$*"
  local end=$((SECONDS + timeout))
  while :; do
    bash -c "$cmd" >/dev/null 2>&1
    rc=$?
    if [ $rc -eq 0 ]; then
      return 0
    fi
    if [ $SECONDS -ge $end ]; then
      echo "Timed out waiting for: $cmd" >&2
      return 1
    fi
    sleep "$sleep_time"
  done
}

archive_info() {
  local out="/tmp/fortanix-test-logs-$(date +%s).tar.gz"
  echo "Collecting debug info to ${out} ..." >&3
  tmpdir=$(mktemp -d)
  {
    kubectl --namespace "$NAMESPACE" get pods -o wide >"$tmpdir/pods.txt" 2>&1 || true
    kubectl --namespace "$NAMESPACE" get secretproviderclasses.secrets-store.csi.x-k8s.io -o yaml >"$tmpdir/spc.yaml" 2>&1 || true
    kubectl --namespace "$NAMESPACE" get events --sort-by='.lastTimestamp' >"$tmpdir/events.txt" 2>&1 || true
    for p in $(kubectl --namespace "$NAMESPACE" get pods -l app=fortanix-csi-provider -o jsonpath='{.items[*].metadata.name}' 2>/dev/null); do
      kubectl --namespace "$NAMESPACE" logs "$p" --all-containers=true >"$tmpdir/logs-${p}.txt" 2>&1 || true
      kubectl --namespace "$NAMESPACE" describe pod "$p" >"$tmpdir/describe-${p}.txt" 2>&1 || true
    done
  } || true
  tar -czf "$out" -C "$tmpdir" . || true
  rm -rf "$tmpdir" || true
  echo "Saved debug archive: $out" >&3
}

# ------------------------
# Per-file setup/teardown
# ------------------------
setup_file() {
  echo "Running setup_file..." >&3

  # Skip early if required envs are not set
  [[ -n "${FORTANIX_API_KEY}" ]] || skip "FORTANIX_API_KEY not set"
  [[ -n "${FORTANIX_DSM_ENDPOINT}" ]] || skip "FORTANIX_DSM_ENDPOINT not set"
  [[ -n "${FORTANIX_TEST_SECRET_NAME}" ]] || skip "FORTANIX_TEST_SECRET_NAME not set"

  # Create the Kubernetes secret with the fixed name the provider expects (fortanix-api-key).
  # README instructs creating "fortanix-api-key" in kube-system with key "api-key". :contentReference[oaicite:2]{index=2}
  cat <<EOF | kubectl --namespace "$NAMESPACE" apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: fortanix-api-key
  namespace: ${NAMESPACE}
type: Opaque
stringData:
  api-key: ${FORTANIX_API_KEY}
EOF

  echo "Created kubernetes secret: fortanix-api-key in ${NAMESPACE}" >&3
}

teardown_file() {
  echo "Running teardown_file..." >&3
  kubectl delete pod "$POD_NAME" --namespace "$NAMESPACE" --ignore-not-found=true || true
  kubectl delete secretproviderclass fortanix-test --namespace "$NAMESPACE" --ignore-not-found=true || true
  kubectl delete secret fortanix-api-key --namespace "$NAMESPACE" --ignore-not-found=true || true
  # optionally remove provider deployment if you want:
  # kubectl delete -f "$PROVIDER_YAML" --namespace "$NAMESPACE" --ignore-not-found=true || true
  archive_info || true
}

# ------------------------
# Tests
# ------------------------

@test "Install Fortanix provider" {
  # apply the provider YAML from the repository (raw URL)
  run kubectl --namespace "$NAMESPACE" apply -f "$PROVIDER_YAML"
  assert_success

  # wait for provider pod(s) to be Ready
  wait_for_process $WAIT_TIME $SLEEP_TIME "kubectl --namespace $NAMESPACE wait --for=condition=Ready --timeout=30s pod -l app=fortanix-csi-provider" || {
    echo "Provider pods did not become ready in time" >&2
    archive_info
    return 1
  }

  PROVIDER_POD=$(kubectl --namespace "$NAMESPACE" get pod -l app=fortanix-csi-provider -o jsonpath="{.items[0].metadata.name}")
  run kubectl --namespace "$NAMESPACE" get pod/"$PROVIDER_POD"
  assert_success
}

@test "deploy fortanix SecretProviderClass" {
  # Deploy a SecretProviderClass referencing the secret name the provider uses
  cat <<EOF | kubectl --namespace "$NAMESPACE" apply -f -
apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: fortanix-test
  namespace: ${NAMESPACE}
spec:
  provider: fortanix-csi-provider
  parameters:
    dsmEndpoint: "${FORTANIX_DSM_ENDPOINT}"
    objects: |
      - secretName: "${FORTANIX_TEST_SECRET_NAME}"
EOF

  # wait for the SPC resource to be visible
  cmd="kubectl --namespace $NAMESPACE get secretproviderclasses.secrets-store.csi.x-k8s.io/fortanix-test -o yaml"
  wait_for_process $WAIT_TIME $SLEEP_TIME "$cmd" || {
    echo "SecretProviderClass not created / visible" >&2
    kubectl --namespace "$NAMESPACE" get secretproviderclasses.secrets-store.csi.x-k8s.io -o wide || true
    return 1
  }
}

@test "CSI inline volume mount test" {
  # create a test pod that mounts the SPC via CSI inline volume
  cat <<EOF | kubectl --namespace "$NAMESPACE" apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: ${POD_NAME}
  namespace: ${NAMESPACE}
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

  # wait for the test pod to be ready
  kubectl --namespace "$NAMESPACE" wait --for=condition=Ready --timeout=120s pod/"$POD_NAME"
  run kubectl --namespace "$NAMESPACE" get pod/"$POD_NAME"
  assert_success
}

@test "Check secret file in pod" {
  run kubectl exec "$POD_NAME" --namespace "$NAMESPACE" -- ls /mnt/secrets
  assert_success
  assert_output_contains "${FORTANIX_TEST_SECRET_NAME}"
  echo "✓ Secret file found: ${FORTANIX_TEST_SECRET_NAME}" >&3
}

@test "Check secret content is non-empty" {
  run kubectl exec "$POD_NAME" --namespace "$NAMESPACE" -- cat "/mnt/secrets/${FORTANIX_TEST_SECRET_NAME}"
  assert_success
  if [ -z "${output:-}" ]; then
    echo "Secret content empty" >&2
    archive_info
    return 1
  fi
  echo "✓ Secret content retrieved" >&3
}

@test "Unmount and delete pod, ensure cleanup" {
  run kubectl --namespace "$NAMESPACE" delete pod/"$POD_NAME" --ignore-not-found=false
  assert_success

  kubectl --namespace "$NAMESPACE" wait --for=delete --timeout=${WAIT_TIME}s pod/"$POD_NAME"
  assert_success

  run kubectl --namespace "$NAMESPACE" get secretproviderclass fortanix-test
  assert_success
}

@test "Cleanup SecretProviderClass and api-key secret" {
  run kubectl --namespace "$NAMESPACE" delete secretproviderclass fortanix-test --ignore-not-found=false
  assert_success

  run kubectl --namespace "$NAMESPACE" delete secret fortanix-api-key --namespace "$NAMESPACE" --ignore-not-found=false
  assert_success
}

# teardown_file() above will run automatically

