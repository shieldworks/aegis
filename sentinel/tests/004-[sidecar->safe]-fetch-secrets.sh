#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# `workload` -> `safe`

# `SecureWorkloadToken` has been shared with the workload by `notary`.
# This shall deliver the token that admin has created for this workload.
http POST http://localhost:8017/v1/fetch \
  secret=NotaryGeneratedRandomWorkloadSecret \
  workload=aegis-demo-workload
