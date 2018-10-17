#!/usr/bin/env bash

set -o errexit
set -o errtrace
set -o nounset
set -o pipefail

NAMESPACE="${1:-krs}"

# make sure the namespace exists before we proceed:
kubectl get ns | grep $NAMESPACE > /dev/null || (echo Aborting e2e test since I can not find namespace $NAMESPACE && exit 1)

# deploy,rc,pods
echo "Launching two deployments"
kubectl -n $NAMESPACE run appserver --image centos:7 -- \
           sh -c "while true; do echo WORK ; sleep 10 ; done"

kubectl -n $NAMESPACE run otherserver --image centos:7 -- \
           sh -c "while true; do echo WORK ; sleep 10 ; done"

# svc
echo "Creating a service"
kubectl -n $NAMESPACE expose deploy/appserver --port 80

sleep 10

echo "Deleting two deployments and the service"
kubectl -n $NAMESPACE delete deploy/appserver deploy/otherserver
kubectl -n $NAMESPACE delete svc/appserver 

# ds

# sts

# jobs

# cj

# hpa

