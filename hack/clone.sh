#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

cd ..

git clone git@github.com:zerotohero-dev/aegis-spire.git
git clone git@github.com:zerotohero-dev/aegis-core.git
git clone git@github.com:zerotohero-dev/aegis-sdk-go.git
git clone git@github.com:zerotohero-dev/aegis-safe.git
git clone git@github.com:zerotohero-dev/aegis-sentinel.git
git clone git@github.com:zerotohero-dev/aegis-sidecar.git
git clone git@github.com:zerotohero-dev/aegis-init-container.git
git clone git@github.com:zerotohero-dev/aegis-workload-demo-using-sidecar.git
git clone git@github.com:zerotohero-dev/aegis-workload-demo-using-sdk.git
git clone git@github.com:zerotohero-dev/aegis-web.git

echo "Everything is awesome!"
