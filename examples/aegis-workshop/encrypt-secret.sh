#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

. ./env.sh

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -s '{"username": "*root*", "password": "*Ca$#C0w*", "value": "!AegisRocks!"}' \
  -e
