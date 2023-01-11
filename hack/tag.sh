#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.7.3

git tag -s v$VERSION
git push origin --tags

cd ../aegis-safe || exit
git tag -s v$VERSION
git push origin --tags
cd ..

cd ../aegis-sentinel || exit
git tag -s v$VERSION
git push origin --tags
cd ../aegis || exit

cd ../aegis-sidecar || exit
git tag -s v$VERSION
git push origin --tags
cd ../aegis || exit

cd ../aegis-workload-demo-using-sidecar || exit
git tag -s v$VERSION
git push origin --tags
cd ../aegis || exit

cd ../aegis-workload-demo-using-sdk || exit
git tag -s v$VERSION
git push origin --tags
cd ../aegis || exit
