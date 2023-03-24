#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

kubectl apply -f ./k8s/sentinel/Namespace.yaml
kubectl apply -f ./k8s/sentinel/ServiceAccount.yaml
kubectl apply -f ./k8s/sentinel/Identity.yaml
kubectl apply -f ./k8s/sentinel/Deployment.yaml
