#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

cd ./examples/multiple-secrets || exit

kubectl apply -f ./k8s/ServiceAccount.yaml
kubectl apply -f ./k8s/Deployment.yaml
kubectl apply -f ./k8s/Identity.yaml
