#!/usr/bin/env bash

set -o errexit
set -o errtrace
set -o nounset
set -o pipefail

echo "Launching two deployments"

kubectl -n krs run appserver --image centos:7 -- \
           sh -c "while true; do echo WORK ; sleep 10 ; done"

kubectl -n krs run otherserver --image centos:7 -- \
           sh -c "while true; do echo WORK ; sleep 10 ; done"

echo "Creating service"
kubectl -n krs expose deploy/appserver --port 80

sleep 10

echo "Deleting two deployments now"

kubectl -n krs delete deploy/appserver deploy/otherserver
kubectl -n krs delete svc/appserver 