#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

INITIAL_IMAGE_COUNT=$(docker images | grep -c aegis)

if [ "$INITIAL_IMAGE_COUNT" -eq 0 ]; then
    echo ""
    echo "There are no Aegis images in the registry."
    echo "Are you sure using the minikube docker?"
    echo ""
    echo "Also make sure you have executed 'eval \"\$(minikube docker-env)\'"
    echo "before building images."
    echo ""
    exit 1
fi

make demo-sdk-local

sleep 10

SECRET=$(openssl rand -base64 16)
SENTINEL_POD_NAME=$(kubectl get po -n aegis-system | grep "aegis-sentinel-" | awk '{print $1}')
WORKLOAD_POD_NAME=$(kubectl get po -n default | grep "aegis-workload-demo-" | awk '{print $1}')

kubectl exec "$SENTINEL_POD_NAME" -n aegis-system -- aegis -w "aegis-workload-demo" -s "$SECRET"

echo "will wait for 30 seconds"
sleep 30

if kubectl logs "$WORKLOAD_POD_NAME" -n default | grep -q "$SECRET"; then
  echo ""
  echo "PASS \o/"
  echo ""
else
  echo ""
  echo "FAIL :("
  echo ""
  exit 1
fi

echo ""
echo "sdk test done… moving onto sidecar test…"
echo ""

make demo-sidecar-local

sleep 10

SECRET=$(openssl rand -base64 16)
SENTINEL_POD_NAME=$(kubectl get po -n aegis-system | grep "aegis-sentinel-" | awk '{print $1}')
WORKLOAD_POD_NAME=$(kubectl get po -n default | grep "aegis-workload-demo-" | awk '{print $1}')

kubectl exec "$SENTINEL_POD_NAME" -n aegis-system -- aegis -w "aegis-workload-demo" -s "$SECRET"

echo "will wait for 30 seconds"
sleep 30

if kubectl logs "$WORKLOAD_POD_NAME" -n default | grep -q "$SECRET"; then
  echo ""
  echo "PASS \o/"
  echo ""
else
  echo ""
  echo "FAIL :("
  echo ""
  exit 1
fi

echo ""
echo "sidecar test is done. moving onto init container test…"
echo ""

make demo-init-container-local

sleep 10

SECRET=$(openssl rand -base64 16)
SENTINEL_POD_NAME=$(kubectl get po -n aegis-system | grep "aegis-sentinel-" | awk '{print $1}')
WORKLOAD_POD_NAME=$(kubectl get po -n default | grep "aegis-workload-demo-" | awk '{print $1}')

kubectl exec "$SENTINEL_POD_NAME" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-n "default" \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
-k

sleep 10

IMAGE_COUNT=$(kubectl get po -n default | grep "$WORKLOAD_POD_NAME" | grep -c Running)

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

echo ""
echo "Everything is awesome!"
echo ""