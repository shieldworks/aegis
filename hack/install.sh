#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

cd ./install/k8s || exit

kubectl apply -k ./spire

while ! kubectl get po -n spire-system | grep spire-server | grep Running
do
    echo "waiting for spire server to be up."
    sleep 5
done

while ! kubectl get po -n spire-system | grep spire-agent | grep Running
do
    echo "waiting for spire agent to be up."
    sleep 2
done

cd safe || exit
# TODO: this will need an update!
kubectl apply -f Namespace.yaml
kubectl apply -f Identity.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -f Deployment.yaml
kubectl apply -f Service.yaml

cd ..
cd sentinel || exit
kubectl apply -f Namespace.yaml
kubectl apply -f Identity.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -f Deployment.yaml

echo "Everything is awesome!"
