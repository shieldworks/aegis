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
  -w "example" \
  -n "default" \
  -f "yaml" \
  -s '{"username": "admin", "password": "CashCow!", "value": "AegisRocks!"}' \
  -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
  -a
