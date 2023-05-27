#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

PACKAGE="$1"
BINS="$2"

echo "Intentionally skipping local build: ${PACKAGE} ${BINS}"
echo "Will use a container image to do the build instead."

# go mod vendor
# go build -o "${PACKAGE}" "${BINS}"
