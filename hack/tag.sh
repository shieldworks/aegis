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
docker trust sign aegishub/aegis-safe:"$VERSION"
echo "aegis-sentinel"
docker trust sign aegishub/aegis-sentinel:"$VERSION"
echo "aegis-sidecar"
docker trust sign aegishub/aegis-sidecar:"$VERSION"
echo "aegis-init-container"
docker trust sign aegishub/aegis-init-container:"$VERSION"
echo "aegis-workload-demo-using-sidecar"
docker trust sign aegishub/aegis-workload-demo-using-sidecar:"$VERSION"
echo "aegis-workload-demo-using-sdk"
docker trust sign aegishub/aegis-workload-demo-using-sdk:"$VERSION"
echo "aegis-workload-demo-using-init-container"
docker trust sign aegishub/aegis-workload-demo-using-init-container:"$VERSION"

echo "aegis-web"
cd ../aegis-web || exit
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo "Everything is awesome!"
