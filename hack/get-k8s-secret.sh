#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

kubectl get secret aegis-secret-aegis-workload-demo -n default \
-o=jsonpath='{.data.VALUE}' | base64 --decode
