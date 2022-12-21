#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

http PUT http://aegis-wokload-demo:8039/v1/hook \
  id=AegisRocks \
  notaryToken=NotaryGeneratedSecureWorkloadToken \
	workloadId=aegis-workload-demo\
	workloadSecret=NotaryGeneratedRandomWorkloadSecret \
	safeApiRoot="http://aegis-safe:8017/"
