#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')
export SENTINEL=$SENTINEL

SAFE=$(kubectl get po -n aegis-system \
  | grep "aegis-safe-" | awk '{print $1}')
export SAFE=$SAFE

