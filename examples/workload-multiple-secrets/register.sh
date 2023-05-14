#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "aegis-workload-demo" \
  -n "default" \
  -s '{"name": "USERNAME", "value": "admin"}' \
  -a

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "aegis-workload-demo" \
  -n "default" \
  -s '{"name": "PASSWORD", "value": "AegisRocks!"}' \
  -a
