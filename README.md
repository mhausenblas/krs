# krs—Kubernetes resource stats

`krs` is a command line tool for capturing and serializing Kubernetes resource statistics in [OpenMetrics](https://github.com/OpenObservability/OpenMetrics) format. It dumps statistics about Kubernetes resources, for example the number of pods in a certain namespace, on a periodic basis to local storage. The kind of resources (pods, services, etc.) as well as the scope, that is, cluster-level or a list of namespaces, is configurable. You can use `krs` either on the client-side (for example, from your laptop) or in-cluster, like in a deployment. Note that `krs` leaves the decision where and how long-term storage is carried out up to you.

## Install

## Use

## Config