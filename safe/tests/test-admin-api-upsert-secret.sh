#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# `admin` -> `safe`

# Keep the admin token safe; do not store it in source control.
# An ideal place to store it is a password manager or an encrypted file.
# Admin retrieves secure admin token via a secure means after `notary`
# bootstraps `safe`.
http PUT http://localhost:8017/v1/secret \
  token=SecureAdminToken \
  key=aegis-demo-workload \
  value='{"username": "me@volkan.io", "password": "ToppyTopSecret"}'
