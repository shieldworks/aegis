#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

. ./env.sh

# Freeform transformation
kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-n "default" \
-f "none" \
-s '{"username": "admin", "password": "cashcow!", "value": "AegisRocks!"}' \
-t 'NEITHER JSON "{{.username}}", "PASSWORD":"{{.password}}" NOR YAML "VALUE" {{.value}}"' \
-a
