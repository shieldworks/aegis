#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#


export SECRET="ComputeMe!"

export SECRET=YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IFgyNTUxOSBBamVPTWZaM3ZOd3RtZzZ4cExVM05LeWVodTlaSm9FVTdQMEFWMko0NEF3CnlWRnFJM0ZxL2ZSVm1yelJCY3YrOWVLUnBud2pjV2pwK3FJVzl4QmRUMFEKLS0tIElBVFVJSXoySDNBcHZEaElIRWJVY1JIOTUwV3FUSzJWMVlpOGxwbVYvbTAKguxim6t6jJRmUlLO8S/4//P/YruBvn4SvgLDSteFd/vLMtZom156kMx0f42jlTK1ZJpxtCbnPGoVfAlTtIW66hIN1VX2k5E2jpUhxRPOyNt930WZNn6p1njYbISHy8jmvmrOQIWg


# export SECRET=YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IFgyNTUxOSBJS0NQTzk0YzRIY1RQc3ZNay94MXAzWlhXVHhxc2FhNSt4VnFlZjZHZ2pVCnlIbUJhb3plSUFBaVo5VVFqWlBobk1uU2dyVklEcGlNYm9XSURtK3YzYmcKLS0tIExrT2RlZTQ4bDNTWThiVUVJUXgxY0lTZCtOV1R6K0pETitUUXFYUmZVZnMKKu2pVaNiL1M+NntkBs/unhuvzVJzqKGffYcGR5Hd59D6VOdrwqF6oRq1Z50vKYwRAmgbwcSGF7itpBetYZpynqa+SkncBezZ/RyrfRK/HOcSv0EFTdYfr13dQupnhHv4wrr0Zm99


export SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')

export SAFE=$(kubectl get po -n aegis-system \
  | grep "aegis-safe-" | awk '{print $1}')

export WORKLOAD=$(kubectl get po -n default \
  | grep "aegis-workload-demo-" | awk '{print $1}')

export INSPECTOR=$(kubectl get po -n default \
  | grep "aegis-inspector-" | awk '{print $1}')

export DEPLOYMENT="aegis-workload-demo"