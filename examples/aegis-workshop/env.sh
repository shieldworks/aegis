#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#


export SECRET="ComputeMe!"

SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')
export SENTINEL=$SENTINEL

SAFE=$(kubectl get po -n aegis-system \
  | grep "aegis-safe-" | awk '{print $1}')
export SAFE=$SAFE

WORKLOAD=$(kubectl get po -n default \
  | grep "aegis-workload-demo-" | awk '{print $1}')
export WORKLOAD=$WORKLOAD

INSPECTOR=$(kubectl get po -n default \
  | grep "aegis-inspector-" | awk '{print $1}')
export INSPECTOR=$INSPECTOR

export DEPLOYMENT="aegis-workload-demo"
