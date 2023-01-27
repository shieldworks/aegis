#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

SENTINEL=aegis-sentinel-58f6478b79-6g242

kubectl exec -it $SENTINEL \
-n aegis-system -- aegis \
-w aegis-safe \
-s '{"logLevel": 1}'
