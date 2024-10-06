# Readiness probe

```shell
k create -f ks-centos-readinessprobe.yaml
```

```shell
k create -f ks-centos-readinessprobe.yaml
# deployment.apps/centos created

k get po
# NAME                      READY   STATUS    RESTARTS   AGE
# centos-7d68b87784-4m4v8   0/1     Running   0          6s
```

```shell
k exec -it centos-7d68b87784-4m4v8 -- /bin/bash

[root@centos-7d68b87784-4m4v8 /] touch /tmp/healthy
[root@centos-7d68b87784-4m4v8 /] exit
# exit

k get po
# NAME                      READY   STATUS    RESTARTS   AGE
# centos-7d68b87784-4m4v8   1/1     Running   0          115s
```