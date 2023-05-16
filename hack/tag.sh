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
echo "aegis-sentinel"
docker trust sign aegishub/aegis-ist-sentinel:"$VERSION"
echo "aegis-sidecar"
docker trust sign aegishub/aegis-ist-sidecar:"$VERSION"
echo "aegis-init-container"
docker trust sign aegishub/aegis-ist-init-container:"$VERSION"
echo "example-using-sidecar"
docker trust sign aegishub/example-using-sidecar:"$VERSION"
echo "example-using-sdk"
docker trust sign aegishub/example-using-sdk:"$VERSION"
echo "example-multiple-secrets"
docker trust sign aegishub/example-multiple-secrets:"$VERSION"
echo "example-using-init-container"
docker trust sign aegishub/example-using-init-container:"$VERSION"

echo "aegis-web"
cd ../aegis-web || exit
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo "Everything is awesome!"
