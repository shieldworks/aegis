#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# Minikube might need additional flags for SPIRE to work properly.
# A bare-metal or cloud Kubernetes cluster will not need these extra configs.
minikube start \
    --extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key \
    --extra-config=apiserver.service-account-key-file=/var/lib/minikube/certs/sa.pub \
    --extra-config=apiserver.service-account-issuer=api \
    --extra-config=apiserver.api-audiences=api,spire-server \
    --extra-config=apiserver.authorization-mode=Node,RBAC \
    --insecure-registry "10.0.0.0/24"

minikube addons enable registry

# If you don't see a registry pod in `docker ps`
# kubectl port-forward --namespace kube-system service/registry 5000:80
