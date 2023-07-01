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

echo "aegis-web"
cd ../aegis-web || exit
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo "Everything is awesome!"
