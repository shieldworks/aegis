#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.7.3

echo "aegis"
git tag -s v$VERSION
git push origin --tags


gh release create

echo "aegis-safe"
cd ../aegis-safe || exit
git tag -s v$VERSION
git push origin --tags
gh release create
docker trust sign z2hdev/aegis-safe:$VERSION
cd ../aegis || exit

echo "aegis-sentinel"
cd ../aegis-sentinel || exit
git tag -s v$VERSION
git push origin --tags
gh release create
docker trust sign z2hdev/aegis-sentinel:$VERSION
cd ../aegis || exit

echo "aegis-sidecar"
cd ../aegis-sidecar || exit
git tag -s v$VERSION
git push origin --tags
gh release createe
docker trust sign z2hdev/aegis-sidecar:$VERSION
cd ../aegis || exit

echo "aegis-workload-demo-using-sidecar"
cd ../aegis-workload-demo-using-sidecar || exit
git tag -s v$VERSION
git push origin --tags
gh release create
docker trust sign z2hdev/aegis-workload-demo-using-sidecar:$VERSION
cd ../aegis || exit

echo "aegis-workload-demo-using-sdk"
cd ../aegis-workload-demo-using-sdk || exit
git tag -s v$VERSION
git push origin --tags
gh release create
docker trust sign z2hdev/aegis-workload-demo-using-sdk:$VERSION
cd ../aegis || exit

echo "aegis-spire"
cd ../aegis-spire || exit
git tag -s v$VERSION
git push origin --tags
gh release create
cd ../aegis || exit

echo "aegis-core"
cd ../aegis-core || exit
git tag -s v$VERSION
git push orign --tags
gh release create
cd ../aegis || exit

echo "aegis-sdk-go"
cd ../aegis-sdk-go || exit
git tag -s v$VERSION
git push origin --tags
gh release create
cd ../aegis || exit

echo "aegis-web"
cd ../aegis-web || exit
git tag -s v$VERSION
git push origin --tags
gh release create
cd ../aegis || exit

echo "Everything is awesome!"
