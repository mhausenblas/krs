 # krs—Kubernetes resource stats
 
[![Go Report Card](https://goreportcard.com/badge/github.com/mhausenblas/krs)](https://goreportcard.com/report/github.com/mhausenblas/krs)
[![container repository on Quay](https://quay.io/repository/mhausenblas/krs/status "Docker Repository on Quay")](https://quay.io/repository/mhausenblas/krs)


<div style="text-align: center; margin-bottom: 50px;">
  <img src="om-k8s.png" width="200px" />
</div>

<p style="color: #e03030;">
NOTE: this is heavy work in progress and we may introduce breaking, non-backwards compatible changes in any version before 1.0 and also, use it at your own risk, ideally only in dev and test environments.
</p>

`krs` is a command line tool for capturing and serializing Kubernetes resource statistics in [OpenMetrics](https://github.com/OpenObservability/OpenMetrics) format. It dumps statistics about Kubernetes resources, for example the number of pods in a certain namespace, on a periodic basis to local storage. The kind of resources (pods, services, etc.) as well as the scope, that is, cluster-level or a list of namespaces, is configurable. You can use `krs` either on the client-side (for example, from your laptop) or in-cluster, like in a deployment. Note that `krs` leaves the decision where and how long-term storage is carried out up to you.

---

Index:

- [Install](#install)
    - [From binaries](#from-binaries)
    - [From source](#from-source)
    - [From Kubernetes](#from-kubernetes)
- [Use](#use)


## Install

In order to use `krs` you must meet the following two prerequisites:

1. `kubectl` must be [installed](https://kubernetes.io/docs/tasks/tools/install-kubectl/).
1. Access to a Kubernetes cluster must be configured. 

Here are my test environments: a v1.9 cluster via OpenShift Online, a v1.10 cluster via AKS, and a v1.11 cluster via Minikube, 
all with client-side with a `kubectl`@v1.11 on macOS.

### From binaries

Binaries for the following platforms are available:

- [Linux](https://github.com/mhausenblas/krs/releases/download/0.1/krs_linux) 
- [macOS](https://github.com/mhausenblas/krs/releases/download/0.1/krs_macos) 
- [Windows](https://github.com/mhausenblas/krs/releases/download/0.1/krs_windows)

To install from binary, for example, on a macOS system, do:

```shell
$ curl -sL https://github.com/mhausenblas/krs/releases/download/0.1/krs_macos -o krs
$ chmod +x krs
$ sudo mv krs /usr/local/bin
```

### From source

Assuming you've got Go in version 1.10 or above installed you can install `krs` from source like so:

```shell
$ go get -u github.com/mhausenblas/krs
```

### From Kubernetes

You can launch `krs` in Kubernetes using the `launch.sh` script as follows. The script will dynamically create the respective permissions for the namespace and run `krs` as a deployment.  
So, in a sense this is a self-test: create and watch a namespace `krs-test` and view the output like so:

```shell
$ kubectl create ns krs-test
namespace/krs-test created

$ ./launch.sh krs-test
serviceaccount/krs created
clusterrole.rbac.authorization.k8s.io/resreader configured
rolebinding.rbac.authorization.k8s.io/allowpodprobes created
deployment.apps/krs created

$ kubectl -n krs-test logs -f $(kubectl -n krs-test get po -l=run=krs --output=jsonpath={.items[*].metadata.name})
# HELP pods Number of pods in any state, for example running
# TYPE pods gauge
pods{namespace="krs-test"} 1
# HELP deployments Number of deployments
# TYPE deployments gauge
deployments{namespace="krs-test"} 1
# HELP services Number of services
# TYPE services gauge
services{namespace="krs-test"} 0
...
```

## Use

`krs` assumes that `kubectl` is installed and configured. It writes the OpenMetrics data to `stdout` which you can then redirect to a file or process further. 

For example, to gathers stats of the `dev42` namespace and store the OpenMetrics formatted  stats in a file called `/tmp/krs/2018-10-05.om` as well as see the errors on screen (via `stdout`), do the following:

```shell
$ krs dev42 >> /tmp/krs/2018-10-05.om
```

If you don't provide a namespace as the first argument, `krs` will watch the `default` namespace. Note that with the environment variable `KRS_KUBECTL_BIN` you can set the `kubectl` to use, which, especially under Windows is required.

For example, the [output](e2e-test/e2e-output.om) of the [end-to-end test](e2e-test/e2e.sh) looks as follows:

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

There are two environment variables that `krs` understands:

- With `KRS_KUBECTL_BIN` you can define which `kubectl` to use.
- With `KRS_VERBOSE` set (for example, to `true` but really any value does) you can get debug-level information, down to what `kubectl` call has been issued.

To do:

- [ ] Extend the e2e testing to all supported resource kinds
- [ ] Improve usage docs (more examples) and inline docs
- [ ] Improve [Make file](https://suva.sh/posts/well-documented-makefiles/)
- [ ] Cut 0.2 release (binaries and container image) 