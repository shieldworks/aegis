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
kubectl apply -f .

