#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

VERSION="$1"

echo ""
echo "--------"
echo "aegis"
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo ""
echo "--------"
echo "aegis-safe"
docker trust sign aegishub/aegis-ist-safe:"$VERSION"
docker trust sign aegishub/aegis-ist-safe:latest
echo "aegis-sentinel"
docker trust sign aegishub/aegis-ist-sentinel:"$VERSION"
docker trust sign aegishub/aegis-ist-sentinel:latest
echo "aegis-sidecar"
docker trust sign aegishub/aegis-ist-sidecar:"$VERSION"
docker trust sign aegishub/aegis-ist-sidecar:latest
echo "aegis-init-container"
docker trust sign aegishub/aegis-ist-init-container:"$VERSION"
docker trust sign aegishub/aegis-ist-init-container:latest
echo "example-using-sidecar"
docker trust sign aegishub/example-using-sidecar:"$VERSION"
docker trust sign aegishub/example-using-sidecar:latest
echo "example-using-sdk"
docker trust sign aegishub/example-using-sdk:"$VERSION"
docker trust sign aegishub/example-using-sdk:latest
echo "example-multiple-secrets"
docker trust sign aegishub/example-multiple-secrets:"$VERSION"
docker trust sign aegishub/example-multiple-secrets:latest
echo "example-using-init-container"
docker trust sign aegishub/example-using-init-container:"$VERSION"
docker trust sign aegishub/example-using-init-container:latest

echo "aegis-photon-safe"
docker trust sign aegishub/aegis-photon-safe:"$VERSION"
docker trust sign aegishub/aegis-photon-safe:latest
echo "aegis-photon-sentinel"
docker trust sign aegishub/aegis-photon-sentinel:"$VERSION"
docker trust sign aegishub/aegis-photon-sentinel:latest
echo "aegis-photon-sidecar"
docker trust sign aegishub/aegis-photon-sidecar:"$VERSION"
docker trust sign aegishub/aegis-photon-sidecar:latest
echo "aegis-photon-init-container"
docker trust sign aegishub/aegis-photon-init-container:"$VERSION"
docker trust sign aegishub/aegis-photon-init-container:latest

echo "aegis-web"
cd ../aegis-web || exit
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo "Everything is awesome!"
