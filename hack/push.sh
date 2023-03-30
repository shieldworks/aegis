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
REPO="$3"

docker build . -t "${PACKAGE}":"${VERSION}"
docker tag "${PACKAGE}":"${VERSION}" "${REPO}":"${VERSION}"
docker push "${REPO}":"${VERSION}"

sleep 10