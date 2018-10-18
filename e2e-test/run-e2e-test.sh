#!/usr/bin/env bash

set -o errexit
set -o errtrace
set -o nounset
set -o pipefail

NAMESPACE="${1:-krs}"

# make sure the namespace exists before we proceed:
kubectl get ns | grep $NAMESPACE > /dev/null || (echo Aborting e2e test since I can not find namespace $NAMESPACE && exit 1)

################################################################################
# base tests (pod,rs,deploy,svc)
# baset $NAMESPACE

###############################################################################
# ds
echo "Creating a daemon set"
kubectl -n $NAMESPACE apply -f ds.yaml
sleep 2
delete -n $NAMESPACE ds krs-test-ds

# sts

# jobs

# cj

# hpa


#### FUNCTIONS #################################################################

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