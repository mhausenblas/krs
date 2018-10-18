#!/usr/bin/env bash

set -o errexit
set -o errtrace
set -o nounset
set -o pipefail

NAMESPACE="${1:-default}"

### Set permissions 

kubectl -n $NAMESPACE create sa krs --dry-run --output=yaml > /tmp/krs-perm.yaml

printf "\n---\n" >> /tmp/krs-perm.yaml

kubectl create clusterrole resreader \
        --verb=get --verb=list \
        --resource=pods --resource=deployments --resource=services \
        --resource=replicationcontroller --resource=daemonsets.apps \
        --resource=deployments.apps --resource=replicasets.apps \
        --resource=statefulsets.apps --resource=horizontalpodautoscalers.autoscaling \
        --resource=jobs.batch --resource=cronjobs.batch \
        --dry-run --output=yaml >> /tmp/krs-perm.yaml

printf "\n---\n" >> /tmp/krs-perm.yaml

kubectl -n $NAMESPACE create rolebinding allowresprobes \
        --clusterrole=resreader \
        --serviceaccount=$NAMESPACE:krs \
        --namespace=$NAMESPACE \
        --dry-run --output=yaml >> /tmp/krs-perm.yaml

kubectl -n $NAMESPACE apply -f /tmp/krs-perm.yaml

### launch tool
kubectl -n $NAMESPACE run krs \
        --image=quay.io/mhausenblas/krs:0.2 \
        --serviceaccount=krs \
        --command -- /app/krs $NAMESPACE

# rm /tmp/krs-perm.yaml