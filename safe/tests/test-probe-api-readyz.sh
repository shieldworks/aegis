#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# (cluster internal)

http http://localhost:8017/readyz

# HTTP/1.1 503 Service Unavailable
# Content-Length: 30
# Content-Type: text/plain; charset=utf-8
# Date: Mon, 19 Dec 2022 07:19:03 GMT
#
# Safe has not bootstrapped yet.

