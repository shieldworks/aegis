#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Demo workload that uses `aegis-init-container`
cd ./install/k8s/demo-workload/using-init-container || exit
kubectl apply -f Identity.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -f Secret.yaml
kubectl apply -k .

