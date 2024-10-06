## Simple pod demo

### Run nginx as webserver

```shell
# We should create deployment by 'create deployment' option
k create deployment nginx --image=nginx:1.27.0 --port=80
```

### Show running pod

```shell
k get po
```

### Expose svc

```shell
# When we create deployment by 'create deployment' option, we use selector 'app=nginx'
k expose deploy nginx --selector app=nginx --port=80 --type=NodePort
```

### Check svc detail

```shell
k get svc
```

### Access service

```shell
curl 192.168.34.101:<nodeport>
```
