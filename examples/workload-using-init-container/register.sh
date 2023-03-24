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

# To make the init container exit successfully and initialize the main
# container of the Pod, execute the following script.
#
# This will create a Kubernetes `Secret` that the `main` container is
# injecting as an environment variable and let the container consume
# that `Secret`.
#
# -n : identifies the namespace of the Kubernetes `Secret`.
# -k : means Aegis will update an associated Kubernetes Secret.
# -t : will be used to transform the fields of the payload.
# -s : is the actual value of the secret.
# -t : is the template transformation used to interpolate the value.

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-n "default" \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
-k
