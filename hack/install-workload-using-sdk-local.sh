#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.ist>
#     .\_/.
#

# Demo workload that directly talks to `aegis-safe` using Aegis Go SDK
cd ./install/k8s/demo-workload/using-sdk || exit
kubectl apply -f Identity.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -k .
