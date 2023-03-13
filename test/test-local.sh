#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

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

echo "sdk test done… moving on to sidecar test…"

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




echo "Everything is awesome!"
