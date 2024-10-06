## Run envoy

```shell
k create configmap envoy-config --from-file=envoy.yaml=ks-envoy-config.yaml
k create -f ks-envoy-deployment.yaml
k expose deploy envoy --selector run=envoy --port=10000 --type=NodePort
```

## Access service

```shell
$ curl <NODE IP Address>:<NodePort>
curl 192.168.34.101:32613
```

## Scale up/down/failover

```shell
k scale deploy <deployment-name> --replicas=<n>
```

```shell
k get po
# NAME                     READY   STATUS    RESTARTS   AGE
# envoy-58bcfd8574-g8j4p   1/1     Running   0          5m35s

k scale deploy envoy --replicas=2
# deployment.apps/envoy scaled

k get po
# NAME                     READY   STATUS    RESTARTS   AGE
# envoy-58bcfd8574-g8j4p   1/1     Running   0          5m55s
# envoy-58bcfd8574-x7hfc   1/1     Running   0          4s
```

## exec

```shell
# configMap 加载成一个文件
k exec -it envoy-58bcfd8574-g8j4p -- /bin/bash

root@envoy-58bcfd8574-g8j4p:/# ls
# bin   dev                   etc   lib    lib64   media  opt   root  sbin  sys  usr
# boot  docker-entrypoint.sh  home  lib32  libx32  mnt    proc  run   srv   tmp  var

root@envoy-58bcfd8574-g8j4p:/# ls /etc/env
# environment  envoy/ 
      
root@envoy-58bcfd8574-g8j4p:/# ls /etc/envoy/
# envoy.yaml
```