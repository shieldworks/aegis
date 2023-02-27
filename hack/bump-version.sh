#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.ist>
#     .\_/.
#

VERSION=0.12.70
NEXT_VERSION=0.12.71

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

cd ../aegis-workload-demo-using-init-container/k8s || exit
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

cd ../aegis-init-container || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
cd ../aegis || exit

cd ./hack || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" tag.sh

echo "Everything is awesome!"
