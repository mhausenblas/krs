#!/usr/bin/env bash

set -o errexit
set -o errtrace
set -o nounset
set -o pipefail

################################################################################
### FUNCTIONS ##################################################################

function baset {
    # deploy,rc,pods
    echo "Launching two deployments"
    kubectl -n $1 run appserver --image centos:7 -- \
            sh -c "while true; do echo WORK ; sleep 10 ; done"

    kubectl -n $1 run otherserver --image centos:7 -- \
            sh -c "while true; do echo WORK ; sleep 10 ; done"

    # svc
    echo "Creating a service"
    kubectl -n $1 expose deploy/appserver --port 80

    sleep 10

    echo "Deleting two deployments and the service"
    kubectl -n $1 delete deploy/appserver deploy/otherserver
    kubectl -n $1 delete svc/appserver 
}

################################################################################
### INIT #######################################################################

### commands line arguments
NAMESPACE="${1:-krs}"

### make sure the namespace exists before we proceed
kubectl get ns | grep $NAMESPACE > /dev/null || (echo Aborting e2e test since I can not find namespace $NAMESPACE && exit 1)

################################################################################
### MAIN #######################################################################

################################################################################
### base tests (pod,rs,deploy,svc)
baset $NAMESPACE

###############################################################################
### daemon set
echo "Creating a daemon set"
kubectl -n $NAMESPACE apply -f ds.yaml
sleep 2
echo "Deleting the daemon set"
kubectl delete -n $NAMESPACE ds krs-test-ds

### sts
echo "Creating a stateful set and service"
kubectl -n $NAMESPACE apply -f sts.yaml
sleep 10
echo "Deleting the stateful set and service"
kubectl delete -n $NAMESPACE svc krs-test-sts-svc
kubectl delete -n $NAMESPACE sts krs-test-sts
sleep 10
echo "Deleting the persistent volume"
kubectl delete -n $NAMESPACE pvc data-krs-test-sts-0

### job

### cron job

### hpa

### ingress
