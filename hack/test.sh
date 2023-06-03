#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Enable strict error checking.
set -euo pipefail

ORIGIN=$1
if [[ -z "$ORIGIN" || "$ORIGIN" != "remote" ]]; then
  ORIGIN="local"
fi

printf "This script assumes that you have a local minikube cluster running,\n"
printf "and you have already installed SPIRE and Aegis.\n"
printf "Also, make sure you have executed 'eval \"\$(minikube docker-env)\'\n"
printf "before running this script.\n"
printf "\n"
read -n 1 -s -r -p "Press any key to proceed…"

# ----- Helper Functions -------------------------------------------------------

### Cleanup and Exit ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_

# Output a success message.
happy_exit() {
  printf "\n"
  printf "Everything is awesome!\n"
  printf "\n"
}

# A k8s problem occurred.
sad_cuddle() {
  local msg
  readonly msg=$1

  if [[ -z "$msg" ]]; then
    printf "Called sad_cuddle() without a message.\n"
    exit 1
  fi

  printf "\n"
  printf "Something went wrong. :(\n"
  printf "%s\n" "$msg"
  printf "\n"
  exit 1
}

# Removes the secret and the demo workload deployment.
cleanup() {
  local sentinel
  readonly sentinel=$(define_sentinel)

  printf "Cleanup…\n"

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -d || sad_cuddle "Cleanup: Failed to delete secret."

  kubectl delete deployment example -n default \
    || sad_cuddle "Cleanup: Failed to delete deployment."

  # Wait for the workload to be gone.
  wait_for_example_workload_deletion &
  wait $!
}

# Deletes the secret associated with the 'example' workload.
delete_secret() {
  local sentinel
  readonly sentinel=$(define_sentinel)

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -d || sad_cuddle "delete_secret: Failed to delete secret."
}

### Definitions ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

# Retrieves the name of the 'example' pod.
define_example_workload() {
  local workload
  readonly workload=$(kubectl get po -n default \
    | grep "example-" | awk '{print $1}'; exit $?)
  if [ $? -ne 0 ]; then
    sad_cuddle "define_example_workload: Failed to define workload."
  fi

  printf "%s" "$workload"
}

# Retrieves the name of the 'aegis-sentinel' pod.
define_sentinel() {
  local sentinel
  readonly sentinel=$(kubectl get po -n aegis-system \
    | grep "aegis-sentinel-" | awk '{print $1}'; exit $?)
  if [ $? -ne 0 ]; then
    sad_cuddle "define_sentinel: Failed to define sentinel."
  fi

  printf "%s" "$sentinel"
}

# Retrieves the name of the 'aegis-safe' pod.
define_safe() {
  local safe
  readonly safe=$(kubectl get po -n aegis-system \
    | grep "aegis-safe-" | awk '{print $1}'; exit $?)
  if [ $? -ne 0 ]; then
    sad_cuddle "define_safe: Failed to define safe."
  fi

  printf "%s" "$safe"
}

### Assertions ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_

# Ensures that the argument is not empty.
assert_exists() {
  local res
  readonly res="$1"

  if [[  -n "$res" ]]; then
    printf "\n"
    printf "FAIL :(\n"
    printf "\n"
    exit 1
  else
    printf "\n"
    printf "PASS \o/\n"
    printf "\n"
  fi
}

# Ensure that the argument equals 1.
assert_only_single_pod() {
  local pod_count
  readonly pod_count=$1

  if [[ "$pod_count" -eq 1 ]]; then
    printf "\n"
    printf "PASS \o/\n"
    printf "\n"
  else
    printf "\n"
    printf "FAIL :(\n"
    printf "\n"
    exit 1
  fi
}

# Ensures that the value and the workload’s secret value are equal.
assert_workload_secret_value() {
  local workload
  local value
  local res

  readonly workload=$(define_example_workload)
  readonly value=$1
  if [[ -z "$workload" || -z "$value" ]]; then
    sad_cuddle "assert_workload_secret_value: Failed to define workload or value."
  fi

  readonly res=$(kubectl exec "$workload" -n default -- ./env; exit $?)
  if [[ $? -ne 0 ]]; then
    sad_cuddle "assert_workload_secret_value: Failed to exec kubectl."
  fi

  if [[ "$res" == "$value" ]]; then
    printf "\n"
    printf "PASS \o/\n"
    printf "\n"
  else
    printf "\n"
    printf "FAIL :(\n"
    printf "\n"
    exit 1
  fi
}

# Ensures that the current workload’s secret is empty.
assert_workload_secret_no_value() {
  local workload
  local res

  readonly workload=$(define_example_workload)
  if [[ -z "$workload" ]]; then
    sad_cuddle "assert_workload_secret_no_value: Failed to define workload."
  fi

  readonly res=$(kubectl exec "$workload" -n default -- ./env; exit $?)
  if [ $? -ne 0 ]; then
    sad_cuddle "assert_workload_secret_no_value: Failed to exec kubectl."
  fi

  if [[ -z "$res" ]]; then
    printf "\n"
    printf "PASS \o/\n"
    printf "\n"
  else
    printf "\n"
    printf "FAIL :(\n"
    printf "\n"
    exit 1
  fi
}

# Ensures if Aegis Sentinel can encrypt a secret.
assert_encrypted_secret() {
  local sentinel
  local value

  readonly sentinel=$(define_sentinel)
  readonly value=$1
  if [[ -z "$sentinel" || -z "$value" ]]; then
    sad_cuddle "assert_encrypted_secret: Failed to define sentinel or value."
  fi

  res=$(kubectl exec "$sentinel" -n aegis-system -- aegis \
    -s "$value" \
    -e; exit $?)
  if [[ $? -ne 0 ]]; then
    sad_cuddle "assert_encrypted_secret: Failed to exec kubectl."
  fi

  assert_exists "$res"
}

# Ensures that the workload is running.
assert_workload_is_running() {
  local workload
  local pod_count

  workload=$(define_example_workload)
  if [[ -z "$workload" ]]; then
    sad_cuddle "assert_workload_is_running: Failed to define workload."
  fi

  pod_count=$(kubectl get po -n default | grep "$workload" | grep -c Running; exit $?)
  if [[ $? -ne 0 ]]; then
    sad_cuddle "assert_workload_is_running: Failed to exec kubectl."
  fi
  if [[ -z "$pod_count" ]]; then
    sad_cuddle "assert_workload_is_running: Empty pod_count"
  fi

  assert_only_single_pod "$pod_count"
}

# Ensures that the init container is running.
assert_init_container_running() {
  local workload
  local pod_status
  readonly workload=$(define_example_workload)

  if [[ -z $workload ]]; then
    sad_cuddle "assert_init_container_running: Failed to define workload."
  fi

  readonly pod_status=$(kubectl get pod -n default "$workload" \
    -o jsonpath='{.status.initContainerStatuses[0].state.running}'; exit $?)
  if [[ $? -ne 0 ]]; then
    sad_cuddle "assert_init_container_running: Failed to exec kubectl."
  fi

  if [[ -n "$pod_status" ]]; then
    printf "Init container of pod '%s' is still running.\n" "$workload"
  else
    printf "Init container of pod '%s' is not running.\n" "$workload"
    exit 1
  fi
}

### Conditions ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_

# Wait for the workload to be ready.
wait_for_example_workload() {
  kubectl wait --for=condition=Ready pod -n default \
    --selector=app.kubernetes.io/name=example || \
    sad_cuddle "wait_for_example_workload: Failed to wait for condition."
}

# Wait until the workload’s deployment is deleted.
wait_for_example_workload_deletion() {
  kubectl wait --for=delete deployment -n default \
    --selector=app.kubernetes.io/name=example || \
    sad_cuddle "wait_for_example_workload_deletion: Failed to wait for deletion."
}

# Pauses the test for 15 seconds to let the sidecar poll the secret.
pause() {
  printf "Waiting for 15 seconds to let the sidecar poll the secret…\n"
  sleep 15
}

# Pauses the test for 15 seconds to let the init container pull the image.
pause_for_deploy() {
  printf "Waiting for 15 seconds to pull the image…\n"
  sleep 15
}

# Pauses the test for 30 seconds to let the init container run, or for other
# operations to complete.
pause_just_in_case() {
  printf "Waiting for 30 seconds, just in case…\n"
  sleep 30
}

### Mutations ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

# Encrypts a secret and stores it in Aegis Safe.
set_encrypted_secret() {
  local value
  local sentinel

  readonly value=$1
  readonly sentinel=$(define_sentinel)
  if [[ -z "$value" || -z "$sentinel" ]]; then
    sad_cuddle "set_encrypted_secret: Failed to define value or sentinel."
  fi

  res=$(kubectl exec "$sentinel" -n aegis-system -- aegis \
    -s "$value" \
    -e; exit $?)
  if [[ $? -ne 0 ]]; then
    sad_cuddle "set_encrypted_secret: Failed to exec kubectl."
  fi
  if [[ -z "$res" ]]; then
    sad_cuddle "set_encrypted_secret: Empty res."
  fi

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s "$res" \
    -e || sad_cuddle "set_encrypted_secret: Failed to exec kubectl."
}

# Registers a secret in Aegis Safe.
set_secret() {
  local sentinel
  local value

  readonly sentinel=$(define_sentinel)
  readonly value=$1
  if [[ -z "$sentinel" || -z "$value" ]]; then
    sad_cuddle "set_secret: Failed to define sentinel or value."
  fi

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s "$value" || sad_cuddle "set_secret: Failed to exec kubectl."
}

# Registers a secret in Aegis Safe and transforms it as JSON.
set_json_secret() {
  local sentinel
  local value
  local transform

  readonly sentinel=$(define_sentinel)
  readonly value=$1
  readonly transform=$2
  if [[ -z "$sentinel" || -z "$value" || -z "$transform" ]]; then
    sad_cuddle "set_json_secret: Failed to define sentinel, value or transform."
  fi

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s "$value" \
    -t "$transform" \
    -f "json" || sad_cuddle "set_json_secret: Failed to exec kubectl."
}

# Registers a secret in Aegis Safe and transforms it as YAML.
set_yaml_secret() {
  local sentinel
  local value
  local transform

  readonly sentinel=$(define_sentinel)
  readonly value=$1
  readonly transform=$2
  if [[ -z "$sentinel" || -z "$value" || -z "$transform" ]]; then
    sad_cuddle "set_yaml_secret: Failed to define sentinel, value or transform."
  fi

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s "$value" \
    -t "$transform" \
    -f "yaml" || sad_cuddle "set_yaml_secret: Failed to exec kubectl."
}

# Registers a secret in Aegis Safe and transforms it as a Kubernetes secret.
set_kubernetes_secret() {
  local sentinel
  readonly sentinel=$(define_sentinel)
  if [[ -z "$sentinel" ]]; then
    sad_cuddle "set_kubernetes_secret: Failed to define sentinel."
  fi

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
    -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "value": "{{.value}}"}' \
    -k || sad_cuddle "set_kubernetes_secret: Failed to exec kubectl."

  # Wait for the workload to be ready.
  wait_for_example_workload &
  wait $!
}

# Append a secret to the workload.
append_secret() {
  local sentinel
  local value

  readonly sentinel=$(define_sentinel)
  readonly value=$1
  if [[ -z "$sentinel" || -z "$value" ]]; then
    sad_cuddle "append_secret: Failed to define sentinel or value."
  fi

  kubectl exec "$sentinel" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -a \
    -s "$value" || sad_cuddle "append_secret: Failed to exec kubectl."
}

### Deployments ### _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

# Deploys a workload that uses Aegis SDK.
deploy_workload_using_sdk() {
  printf "Deploying workload that uses the SDK…\n"

  if [ "$ORIGIN" == "remote" ]; then
    make example-sdk-deploy
  else
    make example-sdk-deploy-local
  fi

  # Wait for the workload to be ready.
  wait_for_example_workload &
  wait $!

  printf "Deployed workload that uses the SDK.\n"
}

# Deploys a workload that uses Aegis Sidecar.
deploy_workload_using_sidecar() {
  printf "Deploying workload that uses the sidecar…\n"

  if [ "$ORIGIN" == "remote" ]; then
    make example-sidecar-deploy
  else
    make example-sidecar-deploy-local
  fi

  # Wait for the workload to be ready.
  wait_for_example_workload &
  wait $!

  printf "Deployed workload that uses the sidecar.\n"
}

# Deploys a workload that uses Aegis Init Container.
deploy_workload_using_init_container() {
  printf "Deploying workload that uses the init container…\n"

  if [ "$ORIGIN" == "remote" ]; then
    make example-init-container-deploy
  else
    make example-init-container-deploy-local
  fi

  pause_for_deploy

  printf "I should have something there.\n"
}

# ------------------------------------------------------------------------------

# Tests the encryption of secrets.
test_encrypting_secrets() {
  printf "Testing: Encrypting secrets…\n"

  local value
  readonly value="!AegisRocks!"

  assert_encrypted_secret $value

  deploy_workload_using_sdk
  define_example_workload

  set_encrypted_secret $value
  assert_workload_secret_value $value

  cleanup

  printf "Tested: Encrypting secrets.\n"
}

test_encrypting_secrets

# ------------------------------------------------------------------------------

cleanup
printf "Case: Workload using Aegis SDK…\n"
deploy_workload_using_sdk

# ------------------------------------------------------------------------------

# Tests the registration of secrets.
test_secret_registration() {
  printf "Testing: Secret registration…\n"

  local value
  readonly value="!AegisRocks!"

  set_secret $value
  assert_workload_secret_value $value

  printf "Tested: Secret registration.\n"
}

test_secret_registration

# ------------------------------------------------------------------------------

# Tests the deletion of secrets.
test_secret_deletion() {
  printf "Testing: Secret deletion…\n"

  delete_secret
  assert_workload_secret_no_value

  printf "Tested: Secret deletion.\n"
}

test_secret_deletion

# ------------------------------------------------------------------------------

# Tests the registration of secrets in append mode.
test_secret_registration_append() {
  printf "Testing: Secret registration (append mode)…\n"

  local secret1
  local secret2
  local value
  readonly secret1="!Aegis"
  readonly secret2="Rocks!"
  readonly value='["'"$secret1"'","'"$secret2"'"]'

  append_secret "$secret1"
  append_secret "$secret2"

  assert_workload_secret_value "$value"
  delete_secret

  printf "Tested: Secret registration (append mode).\n"
}

test_secret_registration_append

# ------------------------------------------------------------------------------

# Tests the registration of secrets in JSON format.
test_secret_registration_json_format() {
  printf "Testing: Secret registration (JSON transformation)…\n"

  local value
  local transform
  readonly value='{"username": "*root*", "password": "*Ca$#C0w*"}'
  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

  set_json_secret "$value" "$transform"

  local transformed
  readonly transformed='{"USERNAME":"*root*", "PASSWORD":"*Ca$#C0w*"}'

  assert_workload_secret_value "$transformed"
  delete_secret

  printf "Tested: Secret registration (JSON transformation).\n"
}

test_secret_registration_json_format

# ------------------------------------------------------------------------------

# Tests the registration of secrets in YAML format.
test_secret_registration_yaml_format() {
  printf "Testing: Secret registration (YAML transformation)…\n"

  local value
  local transform
  readonly value='{"username": "*root*", "password": "*Ca$#C0w*"}'
  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

  set_yaml_secret "$value" "$transform"
  value=$(cat << EOF
  USERNAME: "*root*"
  PASSWORD: "*CashC0w*"
EOF
  )

  assert_workload_secret_value "$value"
  delete_secret

  printf "Tested: Secret registration (YAML transformation).\n"
}

test_secret_registration_yaml_format

# ------------------------------------------------------------------------------

cleanup
printf "Case: Workload using Aegis Sidecar…\n"
deploy_workload_using_sidecar

# Note: for sidecar case, keep in mind that the poll interval is 5 seconds,
# based on the overridden value in the example’s deployment.

# ------------------------------------------------------------------------------

# Tests the registration of secrets using Aegis Sidecar.
test_secret_registration_sidecar() {
  printf "Testing: Secret registration…\n"

  local value
  readonly value="!AegisRocks!"

  set_secret "$value"
  pause
  assert_workload_secret_value "$value"

  printf "Tested: Secret registration.\n"
}

test_secret_registration_sidecar

# ------------------------------------------------------------------------------

# Tests the deletion of secrets using Aegis Sidecar.
test_secret_deletion_sidecar() {
  printf "Testing: Secret deletion…\n"

  delete_secret
  pause
  assert_workload_secret_no_value

  printf "Tested: Secret deletion.\n"
}

test_secret_deletion_sidecar

# ------------------------------------------------------------------------------

# Tests the registration of secrets in append mode using Aegis Sidecar.
test_secret_registration_append_sidecar() {
  printf "Testing Secret registration (append mode)…\n"

  local secret1
  local secret2
  local value
  readonly secret1="!Aegis"
  readonly secret2="Rocks!"
  readonly value='["'"$secret1"'","'"$secret2"'"]'

  append_secret "$secret1"
  append_secret "$secret2"

  pause
  assert_workload_secret_value "$value"
  delete_secret

  printf "Tested: Secret registration (append mode).\n"
}

test_secret_registration_append_sidecar

# ------------------------------------------------------------------------------

# Tests the registration of secrets in JSON format using Aegis Sidecar.
test_secret_registration_json_format_sidecar() {
  printf "Testing Secret registration (JSON transformation)…\n"

  local value
  local transform
  readonly value='{"username": "*root*", "password": "*Ca$#C0w*"}'
  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

  set_json_secret "$value" "$transform"

  local transformed
  readonly transformed='{"USERNAME":"*root*", "PASSWORD":"*Ca$#C0w*"}'

  pause
  assert_workload_secret_value "$transformed"
  delete_secret

  printf "Tested: Secret registration (JSON transformation).\n"
}

test_secret_registration_json_format_sidecar

# ------------------------------------------------------------------------------

# Tests the registration of secrets in YAML format using Aegis Sidecar.
test_secret_registration_yaml_format_sidecar() {
  printf "Testing Secret registration (YAML transformation)…\n"

  local value
  local transform
  readonly value='{"username": "*root*", "password": "*CaShC0w*"}'
  readonly transform='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

  set_yaml_secret "$value" "$transform"

  value=$(cat << EOF
  USERNAME: "*root*"
  PASSWORD: "*CashC0w*"
EOF
  )

  pause
  assert_workload_secret_value "$value"
  delete_secret

  printf "Tested: Secret registration (YAML transformation).\n"
}

test_secret_registration_yaml_format_sidecar

# ------------------------------------------------------------------------------

cleanup
printf "Case: Workload using Aegis Init Container…\n"

# Tests the registration of secrets using Aegis Init Container.
test_init_container() {
  printf "Testing: Init Container…\n"

  deploy_workload_using_init_container
  pause_for_deploy
  define_example_workload

  assert_init_container_running
  pause_just_in_case

  set_kubernetes_secret

  assert_workload_is_running

  printf "Tested: Init Container.\n"
}

# ------------------------------------------------------------------------------

printf "All done. Cleaning up…\n"

cleanup
happy_exit
