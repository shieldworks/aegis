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

# Push version tag
docker tag "${PACKAGE}":"${VERSION}" "${REPO}":"${VERSION}"
docker push "${REPO}":"${VERSION}"

# Push latest tag
docker tag "${PACKAGE}":"${VERSION}" "${REPO}":latest
docker push "${REPO}":latest

sleep 10
