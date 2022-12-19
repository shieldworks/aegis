#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# TODO: move these comments into README instead of keeping them inlined here.

# `notary` -> `safe`

# “Hey `safe`, I am `notary`, here is my id (that you already know)
# and here is the workload token that we will share from this point on when
# communicating with workloads. I will also deliver you workload ids and
# workload secrets. The workloads will identify themselves with those ids
# and secrets.”
# “Ah, by the way, this is the admin token, figure out a way to securely
# dispatch it to the admins. You can store it at say /opt/aegis/admin.token
# for now; and later down the line create mini cli app that will display the
# token only once for the admin to store safely.”
http POST http://localhost:8017/v1/bootstrap \
  id=AegisRocks \
  workloadToken=SecureWorkloadToken \
  adminToken=SecureAdminToken
