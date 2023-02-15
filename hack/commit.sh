#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

cd ..
echo "This will push all of the changes EVERYWHERE to the remote repos."
read -p "Are you sure? " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
  echo "aegis-spire"
	cd aegis-spire                          || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-core"
	cd ../aegis-core                        || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-sdk-go"
	cd ../aegis-sdk-go                      || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-sentinel"
	cd ../aegis-sentinel                    || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-sidecar"
	cd ../aegis-sidecar                     || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-init-container"
	cd ../aegis-init-container              || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-safe"
	cd ../aegis-safe                        || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-web"
	cd ../aegis-web                         || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-workload-demo-using-sidecar"
	cd ../aegis-workload-demo-using-sidecar || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis-workload-demo-using-sdk"
	cd ../aegis-workload-demo-using-sdk     || exit; git add .; git commit -m "bump"; git push origin main;
	echo "aegis"
	cd ../aegis                             || exit; git add .; git commit -m "bump"; git push origin main;
	echo "Everything is awesome!"

fi
