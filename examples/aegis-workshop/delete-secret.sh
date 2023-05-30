#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

. ./env.sh

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -d


















#
#if [ "$ORIGIN" == "remote" ]; then
#  make example-sdk-deploy
#else
#  make example-sdk-deploy-local
#fi
#
#echo "will wait for 10 seconds"
#sleep 10
#
#SECRET=$(openssl rand -base64 16)
#SENTINEL_POD_NAME=$(kubectl get po -n aegis-system | grep "aegis-sentinel-" | awk '{print $1}')
#WORKLOAD_POD_NAME=$(kubectl get po -n default | grep "example-" | awk '{print $1}')
#
#echo "registering secret with Aegis…"
#kubectl exec "$SENTINEL_POD_NAME" -n aegis-system -- aegis -w "example" -s "$SECRET"
#echo "registered secret with Aegis."
#
#echo "will wait for 30 seconds"
#sleep 30
#
#if kubectl logs "$WORKLOAD_POD_NAME" -n default | grep -q "$SECRET"; then
#  echo ""
#  echo "PASS \o/"
#  echo ""
#else
#  echo ""
#  echo "FAIL :("
#  echo ""
#  exit 1
#fi
#
#echo ""
#echo "sdk test done… moving onto sidecar test…"
#echo ""
#
#if [ "$ORIGIN" == "remote" ]; then
#  make example-sidecar-deploy
#else
#  make example-sidecar-deploy-local
#fi
#
#echo "will wait for 10 seconds"
#sleep 10
#
#SECRET=$(openssl rand -base64 16)
#SENTINEL_POD_NAME=$(kubectl get po -n aegis-system | grep "aegis-sentinel-" | awk '{print $1}')
#WORKLOAD_POD_NAME=$(kubectl get po -n default | grep "example-" | awk '{print $1}')
#
#echo "registering secret with Aegis…"
#kubectl exec "$SENTINEL_POD_NAME" -n aegis-system -- aegis -w "example" -s "$SECRET"
#echo "registered secret with Aegis."
#
#echo "will wait for 30 seconds"
#sleep 30
#
#if kubectl logs "$WORKLOAD_POD_NAME" -n default | grep -q "$SECRET"; then
#  echo ""
#  echo "PASS \o/"
#  echo ""
#else
#  echo ""
#  echo "FAIL :("
#  echo ""
#  exit 1
#fi
#
#echo ""
#echo "sidecar test is done. moving onto init container test…"
#echo ""
#
#if [ "$ORIGIN" == "remote" ]; then
#  make example-init-container-deploy
#else
#  make example-init-container-deploy-local
#fi
#
#echo "will wait for 10 seconds"
#sleep 10
#
#SECRET=$(openssl rand -base64 16)
#SENTINEL_POD_NAME=$(kubectl get po -n aegis-system | grep "aegis-sentinel-" | awk '{print $1}')
#WORKLOAD_POD_NAME=$(kubectl get po -n default | grep "example-" | awk '{print $1}')
#
#echo "registering secret with Aegis…"
#kubectl exec "$SENTINEL_POD_NAME" -n aegis-system -- aegis \
#  -w "example" \
#  -n "default" \
#  -s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
#  -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
#  -k
#echo "registered secret with Aegis."
#
#echo "will wait for 30 seconds"
#sleep 30
#
#IMAGE_COUNT=$(kubectl get po -n default | grep "$WORKLOAD_POD_NAME" | grep -c Running)
#
#if [ "$IMAGE_COUNT" -eq 1 ]; then
#  echo ""
#  echo "PASS \o/"
#  echo ""
#else
#  echo ""
#  echo "FAIL :("
#  echo ""
#  exit 1
#fi
#
