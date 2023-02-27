#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.ist>
#     .\_/.
#

if kubectl get ns | grep aegis-system; then
  # Order is important for SPIFFE SCI Driver to properly unmount volumes.
  # ref: https://github.com/spiffe/spiffe-csi#failure-to-terminate-pods-when-driver-is-unhealthy-or-removed
  kubectl delete ns aegis-system
  kubectl delete ns spire-system

  kubectl delete ClusterSPIFFEID aegis-workload-demo
  kubectl delete ClusterSPIFFEID aegis-sentinel
  kubectl delete ClusterSPIFFEID aegis-safe
else
  echo "Nothing to clean."
fi

if kubectl delete deployment aegis-workload-demo -n default; then
  echo "Deleted demo workload too.";
else
  echo "No demo workload to delete?… No worries: That’s fine.";
fi

echo "Everything is awesome!"
