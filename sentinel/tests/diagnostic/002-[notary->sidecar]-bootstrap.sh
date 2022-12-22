#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

http PUT http://aegis-workload-demo.default.svc.cluster.local:8039/v1/hook \
  id=AegisRocks \
  notaryToken=NotaryGeneratedSecureWorkloadToken \
	workloadId=aegis-workload-demo\
	workloadSecret=NotaryGeneratedRandomWorkloadSecret \
	safeApiRoot="http://aegis-safe.aegis-system.svc.cluster.local:8017/"
