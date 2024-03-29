#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

PACKAGE="$1"
VERSION="$2"
DOCKERFILE="$3"

go mod vendor
docker build -f "${DOCKERFILE}" . -t "${PACKAGE}":"${VERSION}"

sleep 10