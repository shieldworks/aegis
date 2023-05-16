#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

source ./env.sh

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -n "default" \
  -s '{"name": "USERNAME", "value": "operator"}' \
  -a

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -n "default" \
  -s '{"name": "PASSWORD", "value": "!KeepYourSecrets!"}' \
  -a
