#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

retval=$(kubectl get po -n default \
  | grep "aegis-workload-demo-" | awk '{print $1}')
export WORKLOAD="$retval"

kubectl logs "$WORKLOAD" -n default -f
