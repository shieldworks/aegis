#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# Demo workload that directly talks to `aegis-safe` using Aegis Go SDK
cd ./install/k8s/demo-workload/using-sdk || exit
kubectl apply -f .