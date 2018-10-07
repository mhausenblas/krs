# krs—Kubernetes resource stats

`krs` is a command line tool for capturing and serializing Kubernetes resource statistics in [OpenMetrics](https://github.com/OpenObservability/OpenMetrics) format. It dumps statistics about Kubernetes resources, for example the number of pods in a certain namespace, on a periodic basis to local storage. The kind of resources (pods, services, etc.) as well as the scope, that is, cluster-level or a list of namespaces, is configurable. You can use `krs` either on the client-side (for example, from your laptop) or in-cluster, like in a deployment. Note that `krs` leaves the decision where and how long-term storage is carried out up to you.

## Install

For the time being, assumes you've got Go v1.10 or above installed and then:

```shell
$ go get -u github.com/mhausenblas/krs
```

Binaries and container image to follow soon.

## Use

`krs` assumes that `kubectl` is installed and configured. It writes the OpenMetrics data to `stdout` which you can then redirect to a file or process further. 

For example, to gathers stats of the `dev42` namespace and store the OpenMetrics formatted  stats in a file called `/tmp/krs/2018-10-05.om` as well as see the errors on screen (via `stdout`), do the following:

```shell
$ krs dev42 >> /tmp/krs/2018-10-05.om
```

If you don't provide a namespace as the first argument, `krs` will watch the `default` namespace.

The beginning of the output of the [end-to-end test](e2e.sh) looks as follows, with the complete output as seen in [e2e-output.om](e2e-output.om):

```
# HELP pods Number of pods in any state, for example running
# TYPE pods gauge
pods{namespace="krs"} 2
# HELP deployments Number of deployments
# TYPE deployments gauge
deployments{namespace="krs"} 2
# HELP services Number of services
# TYPE services gauge
services{namespace="krs"} 1
# HELP pods Number of pods in any state, for example running
# TYPE pods gauge
pods{namespace="krs"} 2
# HELP deployments Number of deployments
# TYPE deployments gauge
deployments{namespace="krs"} 2
# HELP services Number of services
# TYPE services gauge
services{namespace="krs"} 1
```

