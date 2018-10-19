# End-to-end testing `krs`

In one terminal session launch `krs`:

```shell
$ krs --namespace=krs --resources="pods,rs,deploy,ds,sts,pv,pvc,ing"
```

In a second one, launch the end-to-end test script in the `e2e-test` directory:

```shell
$ ./run-e2e-test.sh
```

Optionally, to keep an eye on all resources, in a third session:

```shell
$ watch kubectl -n krs get pods,rs,deploy,ds,sts,pv,pvc,ing
```
