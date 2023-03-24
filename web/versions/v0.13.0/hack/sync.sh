#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

JEKYLL_ENV=production jekyll build

aws s3 sync _site/ s3://aegis.ist/versions/v0.13.0

aws cloudfront create-invalidation --distribution-id EZFGMY32S3BBS --paths "/*"