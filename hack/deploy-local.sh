#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

cd ./k8s || exit

kubectl apply -k ./spire

while ! kubectl get po -n spire-system | grep spire-server | grep Running
do
    echo "waiting for SPIRE server to be up."
    sleep 5
done

while ! kubectl get po -n spire-system | grep spire-agent | grep Running
do
    echo "waiting for SPIRE agent to be up."
    sleep 5
done

cd safe || exit
kubectl apply -f ./Namespace.yaml
kubectl apply -f ./Role.yaml
if kubectl get secret -n aegis-system | grep aegis-safe-age-key; then
  echo "!!! The secret 'aegis-safe-age-key' already exists; not going to override it."
  echo "!!! If you want to modify it, make sure you back it up first."
else
  kubectl apply -f ./Secret.yaml
fi
kubectl apply -f ./ServiceAccount.yaml
kubectl apply -f ./Identity.yaml
kubectl apply -f ./Service.yaml
kubectl apply -k .

cd ..
cd sentinel || exit
kubectl apply -f Namespace.yaml
kubectl apply -f Identity.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -k .

echo "Everything is awesome!"
