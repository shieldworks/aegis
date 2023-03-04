#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

rm -rf ./install/k8s
mkdir -p ./install/k8s
mkdir -p ./install/k8s/demo-workload

cp -rv ../aegis-workload-demo-using-sdk/k8s ./install/k8s/demo-workload/using-sdk
cp -rv ../aegis-workload-demo-using-sidecar/k8s ./install/k8s/demo-workload/using-sidecar
cp -rv ../aegis-workload-demo-using-init-container/k8s ./install/k8s/demo-workload/using-init-container
cp -rv ../aegis-safe/k8s ./install/k8s/safe
cp -rv ../aegis-sentinel/k8s ./install/k8s/sentinel
cp -rv ../aegis-spire/k8s ./install/k8s/spire