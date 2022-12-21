#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

http PUT http://aegis-safe:8017/v1/secret \
  token=NotaryGeneratedSecureAdminToken \
  key=aegis-workload-demo \
  value='{"username": "me@volkan.io", "password": "ToppyTopSecret"}'
