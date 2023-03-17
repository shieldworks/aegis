#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

cd ..

git clone git@github.com:shieldworks/aegis-spire.git
git clone git@github.com:shieldworkds/aegis-core.git
git clone git@github.com:shieldworks/aegis-sdk-go.git
git clone git@github.com:shieldworks/aegis-safe.git
git clone git@github.com:shieldworks/aegis-sentinel.git
git clone git@github.com:shieldworks/aegis-sidecar.git
git clone git@github.com:shieldworks/aegis-init-container.git
git clone git@github.com:shieldworks/aegis-workload-demo-using-sidecar.git
git clone git@github.com:shieldworks/aegis-workload-demo-using-sdk.git
git clone git@github.com:shieldworks/aegis-workload-demo-using-init-container.git
git clone git@github.com:shieldworks/aegis-web.git

echo "Everything is awesome!"
