#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

OLD=aegis-sentinel-58f6478b79-6g242
NEW=aegis-sentinel-58f6478b79-6g242

sed -i "s/$OLD/$NEW/" ./*.sh
