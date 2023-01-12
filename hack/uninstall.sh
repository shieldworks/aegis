#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

if kubectl get ns | grep aegis-system; then
  kubectl delete ns spire-system
	kubectl delete ns aegis-system
	kubectl delete ClusterSPIFFEID aegis-workload-demo
	kubectl delete ClusterSPIFFEID aegis-sentinel
	kubectl delete ClusterSPIFFEID aegis-safe

	echo "Everything is awesome!"
else
  echo "Nothing to clean."
fi

if kubectl delete deployment aegis-workload-demo -n default; then
  echo "Deleted demo workload too.";
else
  echo "No demo workload to delete?… No worries: That’s fine.";
fi

echo "Everything is awesome!"