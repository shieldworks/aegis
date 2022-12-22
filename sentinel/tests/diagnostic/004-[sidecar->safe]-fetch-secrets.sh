#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

http POST http://aegis-safe.aegis-system.svc.cluster.local:8017/v1/fetch \
  secret=NotaryGeneratedRandomWorkloadSecret \
  workload=aegis-workload-demo
