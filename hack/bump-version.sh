#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.7.6
NEXT_VERSION=0.7.7

cd ../aegis-workload-demo-using-sidecar/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ../aegis || exit

cd ../aegis-workload-demo-using-sdk/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ../aegis || exit

cd ../aegis-safe/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ../aegis || exit

cd ../aegis-sentinel/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ../aegis || exit

cd ../aegis-sidecar || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ../aegis || exit

cd ./hack || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" tag.sh

echo "Everything is awesome!"