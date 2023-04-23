#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

. ./env.sh

cd ./workload-init-container || exit

kubectl apply -f ServiceAccount.yaml
kubectl apply -f Secret.yaml
kubectl apply -k .
