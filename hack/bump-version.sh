#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.6.0
NEXT_VERSION=0.6.1

cd ./demo/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ..

cd ./safe/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ..

cd ./sentinel/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ..

cd ./sidecar || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ..

cd ./hack || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" tag.sh

echo "Everything is awesome!"