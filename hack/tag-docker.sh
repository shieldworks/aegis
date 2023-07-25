#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# This is a script to fix unsigned images.
# Normally, signing and pushing should be a single step
# and we should not need to pull the images and sign them again.
# So we’d rarely (if ever) need to use this script.

VERSION="0.18.2"

export DOCKER_CONTENT_TRUST=0

docker pull aegishub/aegis-ist-safe:"$VERSION"
docker pull aegishub/aegis-ist-safe:latest
docker pull aegishub/aegis-ist-sentinel:"$VERSION"
docker pull aegishub/aegis-ist-sentinel:latest
docker pull aegishub/aegis-ist-sidecar:"$VERSION"
docker pull aegishub/aegis-ist-sidecar:latest
docker pull aegishub/aegis-ist-init-container:"$VERSION"
docker pull aegishub/aegis-ist-init-container:latest
docker pull aegishub/example-using-sidecar:"$VERSION"
docker pull aegishub/example-using-sidecar:latest
docker pull aegishub/example-using-sdk:"$VERSION"
docker pull aegishub/example-using-sdk:latest
docker pull aegishub/example-multiple-secrets:"$VERSION"
docker pull aegishub/example-multiple-secrets:latest
docker pull aegishub/example-using-init-container:"$VERSION"
docker pull aegishub/example-using-init-container:latest

docker pull aegishub/aegis-photon-safe:"$VERSION"
docker pull aegishub/aegis-photon-safe:latest
docker pull aegishub/aegis-photon-sentinel:"$VERSION"
docker pull aegishub/aegis-photon-sentinel:latest
docker pull aegishub/aegis-photon-sidecar:"$VERSION"
docker pull aegishub/aegis-photon-sidecar:latest
docker pull aegishub/aegis-photon-init-container:"$VERSION"
docker pull aegishub/aegis-photon-init-container:latest

export DOCKER_CONTENT_TRUST=1

docker trust sign aegishub/aegis-ist-safe:"$VERSION"
docker trust sign aegishub/aegis-ist-safe:latest
docker trust sign aegishub/aegis-ist-sentinel:"$VERSION"
docker trust sign aegishub/aegis-ist-sentinel:latest
docker trust sign aegishub/aegis-ist-sidecar:"$VERSION"
docker trust sign aegishub/aegis-ist-sidecar:latest
docker trust sign aegishub/aegis-ist-init-container:"$VERSION"
docker trust sign aegishub/aegis-ist-init-container:latest
docker trust sign aegishub/example-using-sidecar:"$VERSION"
docker trust sign aegishub/example-using-sidecar:latest
docker trust sign aegishub/example-using-sdk:"$VERSION"
docker trust sign aegishub/example-using-sdk:latest
docker trust sign aegishub/example-multiple-secrets:"$VERSION"
docker trust sign aegishub/example-multiple-secrets:latest
docker trust sign aegishub/example-using-init-container:"$VERSION"
docker trust sign aegishub/example-using-init-container:latest

docker trust sign aegishub/aegis-photon-safe:"$VERSION"
docker trust sign aegishub/aegis-photon-safe:latest
docker trust sign aegishub/aegis-photon-sentinel:"$VERSION"
docker trust sign aegishub/aegis-photon-sentinel:latest
docker trust sign aegishub/aegis-photon-sidecar:"$VERSION"
docker trust sign aegishub/aegis-photon-sidecar:latest
docker trust sign aegishub/aegis-photon-init-container:"$VERSION"
docker trust sign aegishub/aegis-photon-init-container:latest

