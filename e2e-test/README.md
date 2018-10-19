# End-to-end testing `krs`

## Build krs
```shell
make gbuild
```

## Test krs
In one terminal session launch `krs`:

```shell
$ # Set one of the following variables
$ OS=macos
$ # OS=linux
$ # OS=windows

$ # Launch krs
$ ./out/krs_$OS --namespace=krs --resources="pods,rs,deploy,ds,sts,pv,pvc,hpa,ing"
```

In a second one, launch the end-to-end test script in the `e2e-test` directory:

```shell
$ cd e2e-test
$ ./run-e2e-test.sh
```

Optionally, to keep an eye on all resources, in a third session:

```shell
$ watch kubectl -n krs get pods,rs,deploy,ds,sts,pv,pvc,ing
```
