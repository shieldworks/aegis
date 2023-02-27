#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.ist>
#     .\_/.
#

echo "aegis"
git status -s
cd ..

echo "aegis-core"
cd aegis-core || exit; git status -s
cd ..

echo "aegis-safe"
cd aegis-safe || exit; git status -s
cd ..

echo "aegis-sdk-go"
cd aegis-sdk-go || exit; git status -s
cd ..

echo "aegis-sentinel"
cd aegis-sentinel || exit; git status -s
cd ..

echo "aegis-sidecar"
cd aegis-sidecar || exit; git status -s
cd ..

echo "aegis-init-container"
cd aegis-init-container || exit; git status -s
cd ..

echo "aegis-spire"
cd aegis-spire || exit; git status -s
cd ..

echo "aegis-web"
cd aegis-web || exit; git status -s
cd ..

echo "aegis-workload-demo-using-sdk"
cd aegis-workload-demo-using-sdk || exit; git status -s
cd ..

echo "aegis-workload-demo-using-sidecar"
cd aegis-workload-demo-using-sidecar || exit; git status -s
cd ..

echo "aegis-workload-demo-using-init-container"
cd aegis-workload-demo-using-init-container || exit; git status -s
cd ..
