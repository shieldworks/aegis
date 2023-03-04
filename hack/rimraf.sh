#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

cd ..
echo "This will delete all Aegis repos on your workspace irreversibly."
read -p "Are you sure? " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
  rm -rf aegis-spire
  rm -rf aegis-core
  rm -rf aegis-sdk-go
  rm -rf aegis-safe
  rm -rf aegis-sentinel
  rm -rf aegis-sidecar
  rm -rf aegis-workload-demo-using-sidecar
  rm -rf aegis-workload-demo-using-sdk
  rm -rf aegis-web
  echo "Everything is awesome!"
fi
