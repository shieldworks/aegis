#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

cd ..
echo "This will stash all your uncommitted changes EVERYWHERE."
read -p "Are you sure? " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
  echo "aegis-spire"
  mkdir -p aegis-spire
	cd aegis-spire || exit; git stash; git checkout main; git pull;
	echo "aegis-core"
	mkdir -p ../aegis-core
	cd ../aegis-core || exit; git stash; git checkout main; git pull;
	echo "aegis-sdk-go"
	mkdir -p ../aegis-sdk-go
	cd ../aegis-sdk-go || exit; git stash; git checkout main; git pull;
	echo "aegis-sentinel"
	mkdir -p ../aegis-sentinel
	cd ../aegis-sentinel || exit; git stash; git checkout main; git pull;
	echo "aegis-sidecar"
	mkdir -p ../aegis-sidecar
	cd ../aegis-sidecar || exit; git stash; git checkout main; git pull;
	echo "aegis-safe"
	mkdir -p ../aegis-safe
	cd ../aegis-safe || exit; git stash; git checkout main; git pull;
	echo "aegis-web"
	mkdir -p ../aegis-web
	cd ../aegis-web || exit; git stash; git checkout main; git pull;
	echo "aegis-workload-demo-using-sidecar"
	mkdir -p ../aegis-workload-demo-using-sidecar
	cd ../aegis-workload-demo-using-sidecar || exit; git stash; git checkout main; git pull;
	echo "aegis-workload-demo-using-sdk"
	mkdir -p ../aegis-workload-demo-using-sdk
	cd ../aegis-workload-demo-using-sdk || exit; git stash; git checkout main; git pull;
	echo "Everything is awesome!"
fi
