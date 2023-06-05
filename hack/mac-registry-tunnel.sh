#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Use defaults
export DOCKER_TLS_VERIFY=""
export DOCKER_HOST=""
export DOCKER_CERT_PATH=""
export MINIKUBE_ACTIVE_DOCKERD=""

# Get the output of docker ps.
output=$(docker ps)

# Extract the line containing port 5000.
line=$(echo "$output" | grep -F '5000/tcp')

# Extract the port number.
port=$(echo "$line" | perl -nle 'print $1 if m{(\d+)->5000/tcp}')

echo "Found port '$port'!"

if [[ -z "$port" ]]; then
  echo "Could not find port! Exiting!"
  exit 1
fi

# Run socat to forward traffic from port 5000 to the extracted port.
socat TCP-LISTEN:5000,fork,reuseaddr TCP:localhost:"$port"
