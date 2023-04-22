#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

. /home/v0lk4n/Desktop/AEGIS/aegis/examples/aegis-workshop/env.sh

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-e
