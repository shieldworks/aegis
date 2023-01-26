#!/usr/bin/env bash

SENTINEL=aegis-sentinel-58f6478b79-6g242

kubectl exec -it $SENTINEL \
-n aegis-system -- aegis \
-w aegis-workload-demo \
-s '{"username": "pepsi", "password": "cola"}'
