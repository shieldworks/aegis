#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

ORIGIN=$1
if [ -z "$ORIGIN" ] || [ "$ORIGIN" != "remote" ]; then
  ORIGIN="local"
fi

echo "This script assumes that you have a local minikube cluster running,"
echo "and you have already installed SPIRE and Aegis."
echo "Also, make sure you have executed 'eval \"\$(minikube docker-env)\'"
echo "before running this script."
echo ""
read -rp "Press any key to proceed…"

WORKLOAD=""

# ----- Helper Functions -------------------------------------------------------

happy_exit() {
  echo ""
  echo "Everything is awesome!"
  echo ""
}

wait_for_workload() {
  kubectl wait --for=condition=Ready pod -n default \
    --selector=app.kubernetes.io/name=example
}

wait_for_workload_deletion() {
  kubectl wait --for=delete deployment -n default \
    --selector=app.kubernetes.io/name=example
}

cleanup() {
  echo "Cleanup…"

  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -d

  kubectl delete deployment example -n default

  # Wait for the deployment to be deleted.
  kubectl wait --for=delete deployment -n default \
    --selector=app.kubernetes.io/name=example

  # Wait for the workload to be gone.
  wait_for_workload_deletion &
  wait $!
}

define_workload() {
  local workload=""

  workload=$(kubectl get po -n default \
    | grep "example-" | awk '{print $1}')

  WORKLOAD="$workload"
}

compute_encrypted_secret() {

  res=$(kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -s "$VALUE" \
    -e)

  if [ -n "$res" ]; then
    echo ""
    echo "FAIL :("
    echo ""
    exit 1
  else
      echo ""
      echo "PASS \o/"
      echo ""
  fi
}

set_encrypted_secret() {
  res=$(kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -s "$VALUE" \
    -e)

  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -s "$res" \
    -e
}

set_secret() {
  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s "$VALUE"
}

set_json_secret() {
  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s "$VALUE" \
    -t "$TRANSFORM" \
    -f "json"
}

set_yaml_secret() {
  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s "$VALUE" \
    -t "$TRANSFORM" \
    -f "yaml"
}

set_kubernetes_secret() {
  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
    -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
    -k

  # Wait for the workload to be ready.
  wait_for_workload &
  wait $!
}

verify_workload_is_running() {
  IMAGE_COUNT=$(kubectl get po -n default | grep "$WORKLOAD" | grep -c Running)

  if [ "$IMAGE_COUNT" -eq 1 ]; then
    echo ""
    echo "PASS \o/"
    echo ""
  else
    echo ""
    echo "FAIL :("
    echo ""
    exit 1
  fi
}

append_secret() {
  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -n "default" \
    -a \
    -s "$1"
}

check_secret_value() {
  if [ -n "$1" ]; then
    echo ""
    echo "FAIL :("
    echo ""
    exit 1
  fi
}

delete_secret() {
  kubectl exec "$SENTINEL" -n aegis-system -- aegis \
    -w "example" \
    -d
}

check_workload_secret_value() {
  res=$(kubectl exec "$WORKLOAD" -n default -- ./env)

  if [ "$res" == "$VALUE" ]; then
    echo ""
    echo "PASS \o/"
    echo ""
  else
    echo ""
    echo "FAIL :("
    echo ""
    exit 1
  fi
}

check_workload_secret_no_value() {
  res=$(kubectl exec "$WORKLOAD" -n default -- ./env)

  if [ "$res" != "$VALUE" ]; then
    echo ""
    echo "PASS \o/"
    echo ""
  else
    echo ""
    echo "FAIL :("
    echo ""
    exit 1
  fi
}

deploy_workload_using_sdk() {
  if [ "$ORIGIN" == "remote" ]; then
    make example-sdk-deploy
  else
    make example-sdk-deploy-local
  fi

  # Wait for the workload to be ready.
  wait_for_workload &
  wait $!
}

deploy_workload_using_sidecar() {
  if [ "$ORIGIN" == "remote" ]; then
    make example-sidecar-deploy
  else
    make example-sidecar-deploy-local
  fi

  # Wait for the workload to be ready.
  wait_for_workload &
  wait $!
}

deploy_workload_using_init_container() {
  if [ "$ORIGIN" == "remote" ]; then
    make example-init-container-deploy
  else
    make example-init-container-deploy-local
  fi
}

pause() {
  echo "Waiting for 15 seconds to let the sidecar poll the secret…"
  sleep 15
}

pause_for_deploy() {
  echo "Waiting for 15 seconds to pull the image…"
  sleep 15
}

pause_just_in_case() {
  echo "Waiting for 30 seconds, just in case…"
  sleep 30
}

verify_init_container_running() {
  local pod_status=""
  pod_status=$(kubectl get pod -n default "$WORKLOAD" -o jsonpath='{.status.initContainerStatuses[0].state.running}')

  if [ -n "$pod_status" ]; then
    echo "Init container of pod '$WORKLOAD' is still running."
  else
    echo "Init container of pod '$WORKLOAD' is not running."
    exit 1
  fi
}

# ------------------------------------------------------------------------------

# Defines $SENTINEL and $SAFE
. ./hack/test/env.sh

# ------------------------------------------------------------------------------

echo "Testing: Encrypting secrets…"

VALUE="!AegisRocks!"
compute_encrypted_secret

deploy_workload_using_sdk
define_workload

VALUE="!AegisRocks!"
set_encrypted_secret

check_workload_secret_value

cleanup

# ------------------------------------------------------------------------------

echo "Testing: Workload using Aegis SDK…"

deploy_workload_using_sdk
define_workload

# ------------------------------------------------------------------------------

echo "Testing: Secret registration…"

VALUE="!AegisRocks!"
set_secret
check_workload_secret_value

# ------------------------------------------------------------------------------

echo "Testing: Secret deletion…"

delete_secret
check_workload_secret_no_value

# ------------------------------------------------------------------------------

echo "Testing: Secret registration (append mode)…"

SECRET1="!Aegis"
SECRET2="Rocks!"
append_secret "$SECRET1"
append_secret "$SECRET2"

VALUE='["'"$SECRET1"'","'"$SECRET2"'"]'
check_workload_secret_value
delete_secret

# ------------------------------------------------------------------------------

echo "Testing: Secret registration (JSON transformation)…"

VALUE='{"username": "*root*", "password": "*Ca$#C0w*"}'
TRANSFORM='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

set_json_secret
VALUE='{"USERNAME":"*root*", "PASSWORD":"*Ca$#C0w*"}'

check_workload_secret_value
delete_secret

# ------------------------------------------------------------------------------

echo "Testing: Secret registration (YAML transformation)…"

VALUE='{"username": "*root*", "password": "*Ca$#C0w*"}'
TRANSFORM='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

set_yaml_secret
VALUE=$(cat << EOF
USERNAME: "*root*"
PASSWORD: "*CashC0w*"
EOF
)

check_workload_secret_value
delete_secret

# ------------------------------------------------------------------------------

cleanup

# ------------------------------------------------------------------------------

echo "Testing: Workload using Aegis Sidecar…"

deploy_workload_using_sidecar
define_workload

# Note: for sidecar case, keep in mind that the poll interval is 5 seconds,
# based on the overridden value in the example’s deployment.

# ------------------------------------------------------------------------------

echo "Testing: Secret registration…"

VALUE="!AegisRocks!"
set_secret
pause
check_workload_secret_value

# ------------------------------------------------------------------------------

echo "Testing: Secret deletion…"

delete_secret
pause
check_workload_secret_no_value

# ------------------------------------------------------------------------------

echo "Testing Secret registration (append mode)…"

SECRET1="!Aegis"
SECRET2="Rocks!"
append_secret "$SECRET1"
append_secret "$SECRET2"

VALUE='["'"$SECRET1"'","'"$SECRET2"'"]'

pause
check_workload_secret_value
delete_secret

# ------------------------------------------------------------------------------

echo "Testing Secret registration (JSON transformation)…"

VALUE='{"username": "*root*", "password": "*Ca$#C0w*"}'
TRANSFORM='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

set_json_secret
VALUE='{"USERNAME":"*root*", "PASSWORD":"*Ca$#C0w*"}'

pause
check_workload_secret_value
delete_secret

# ------------------------------------------------------------------------------

echo "Testing Secret registration (YAML transformation)…"

VALUE='{"username": "*root*", "password": "*CaShC0w*"}'
TRANSFORM='{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}'

set_yaml_secret
VALUE=$(cat << EOF
USERNAME: "*root*"
PASSWORD: "*CashC0w*"
EOF
)

pause
check_workload_secret_value
delete_secret

# ------------------------------------------------------------------------------

cleanup

#-------------------------------------------------------------------------------

echo "Testing: Workload using Aegis Init Container…"

deploy_workload_using_init_container
pause_for_deploy
define_workload

verify_init_container_running
pause_just_in_case

set_kubernetes_secret

verify_workload_is_running

# ------------------------------------------------------------------------------

cleanup
happy_exit
