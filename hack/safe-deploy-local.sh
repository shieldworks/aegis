#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

kubectl apply -f ./k8s/safe/Namespace.yaml
kubectl apply -f ./k8s/safe/Role.yaml

if kubectl get secret -n aegis-system | grep aegis-safe-age-key; then
  echo "!!! The secret 'aegis-safe-age-key' already exists; not going to override it."
  echo "!!! If you want to modify it, make sure you back it up first."
else
  kubectl apply -f ./k8s/safe/Secret.yaml
fi

kubectl apply -f ./k8s/safe/ServiceAccount.yaml
kubectl apply -f ./k8s/safe.Identity.yaml
kubectl apply -k ./k8s/safe
kubectl apply -f ./k8s/safe/Service.yaml
