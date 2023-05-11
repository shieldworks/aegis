#!/usr/bin/env bash

export SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')

export SAFE=$(kubectl get po -n aegis-system \
  | grep "aegis-safe-" | awk '{print $1}')

export WORKLOAD=$(kubectl get po -n default \
  | grep "aegis-workload-demo-" | awk '{print $1}')

export INSPECTOR=$(kubectl get po -n default \
  | grep "aegis-inspector-" | awk '{print $1}')

export DEPLOYMENT="aegis-workload-demo"

kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-n "default" \
-s "AegisRocks!"
