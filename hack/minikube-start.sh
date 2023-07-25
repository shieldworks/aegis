#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Get CPU and Memory from environment variables or default
CPU_COUNT="${AEGIS_MINIKUBE_CPU_COUNT:-8}"
MEMORY="${AEGIS_MINIKUBE_MEMORY:-11264m}"

# Minikube might need additional flags for SPIRE to work properly.
# A bare-metal or cloud Kubernetes cluster will not need these extra configs.
minikube start \
    --extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key \
    --extra-config=apiserver.service-account-key-file=/var/lib/minikube/certs/sa.pub \
    --extra-config=apiserver.service-account-issuer=api \
    --extra-config=apiserver.api-audiences=api,spire-server \
    --extra-config=apiserver.authorization-mode=Node,RBAC \
    --memory="$MEMORY" \
    --cpus="$CPU_COUNT" \
    --insecure-registry "10.0.0.0/24"

echo "waiting 10 secs before enabling registry"
sleep 10
minikube addons enable registry
kubectl get node
