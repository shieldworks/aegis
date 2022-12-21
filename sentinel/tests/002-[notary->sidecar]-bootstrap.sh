#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# `notary` -> `sidecar`

http PUT http://localhost:8039/v1/hook \
  id=AegisRocks \
  notaryToken=NotaryGeneratedSecureWorkloadToken \
	workloadId=aegis-demo-workload\
	workloadSecret=NotaryGeneratedRandomWorkloadSecret \
	safeApiRoot="http://localhost:8017/"
