#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

cd ./examples/multiple-secrets || exit

kubectl apply -f ./k8s/ServiceAccount.yaml
kubectl apply -f ./k8s/Deployment.yaml
kubectl apply -f ./k8s/Identity.yaml
