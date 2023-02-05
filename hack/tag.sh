#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.12.0

echo ""
echo "--------"
echo "aegis"
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
fi

echo ""
echo "--------"
echo "aegis-safe"
cd ../aegis-safe || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
  docker trust sign z2hdev/aegis-safe:$VERSION
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-sentinel"
cd ../aegis-sentinel || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
  docker trust sign z2hdev/aegis-sentinel:$VERSION
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-sidecar"
cd ../aegis-sidecar || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
  docker trust sign z2hdev/aegis-sidecar:$VERSION
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-workload-demo-using-sidecar"
cd ../aegis-workload-demo-using-sidecar || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
  docker trust sign z2hdev/aegis-workload-demo-using-sidecar:$VERSION
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-workload-demo-using-sdk"
cd ../aegis-workload-demo-using-sdk || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
  docker trust sign z2hdev/aegis-workload-demo-using-sdk:$VERSION
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-spire"
cd ../aegis-spire || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-core"
cd ../aegis-core || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-sdk-go"
cd ../aegis-sdk-go || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
fi
cd ../aegis || exit

echo ""
echo "--------"
echo "aegis-web"
cd ../aegis-web || exit
if git tag -s v$VERSION; then
  git push origin --tags
  gh release create
fi
cd ../aegis || exit

echo "Everything is awesome!"
