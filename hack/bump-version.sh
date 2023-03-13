#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

VERSION=0.13.0
NEXT_VERSION=0.13.1

cd ../aegis-workload-demo-using-sidecar/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
sed -i "s/=\"version=\"$VERSION\"/=\"version=\"$NEXT_VERSION\"/" Dockerfile
cd ../aegis || exit

cd ../aegis-workload-demo-using-sdk/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
sed -i "s/=\"version=\"$VERSION\"/=\"version=\"$NEXT_VERSION\"/" Dockerfile
cd ../aegis || exit

cd ../aegis-workload-demo-using-init-container/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
sed -i "s/=\"version=\"$VERSION\"/=\"version=\"$NEXT_VERSION\"/" Dockerfile
cd ../aegis || exit

cd ../aegis-safe/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
sed -i "s/=\"version=\"$VERSION\"/=\"version=\"$NEXT_VERSION\"/" Dockerfile
cd ../aegis || exit

cd ../aegis-sentinel/k8s || exit
sed -i "s/:$VERSION/:$NEXT_VERSION/" ./*.yaml
cd ..
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
sed -i "s/=\"version=\"$VERSION\"/=\"version=\"$NEXT_VERSION\"/" Dockerfile
cd ../aegis || exit

cd ../aegis-sidecar || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
sed -i "s/=\"version=\"$VERSION\"/=\"version=\"$NEXT_VERSION\"/" Dockerfile
cd ../aegis || exit

cd ../aegis-init-container || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" Makefile
sed -i "s/=\"version=\"$VERSION\"/=\"version=\"$NEXT_VERSION\"/" Dockerfile
cd ../aegis || exit

cd ./hack || exit
sed -i "s/=$VERSION/=$NEXT_VERSION/" tag.sh

echo "Everything is awesome!"
