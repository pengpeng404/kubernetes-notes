
### Start authservice
```shell
# windows
$env:CGO_ENABLED=0
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o bin/amd64 .
```

```shell
# run service
chmod +x bin/amd64/authn-webhook
./bin/amd64/authn-webhook
```

### Create webhook config

```shell
mkdir -p /etc/config
cp webhook-config.json /etc/config
```

### Backup old apiserver
```shell
# 把之前的备份覆盖了
cp /etc/kubernetes/manifests/kube-apiserver.yaml ~/kube-apiserver.yaml
```


### Update apiserver configuration to enable webhook

```shell
cp specs/kube-apiserver.yaml /etc/kubernetes/manifests/kube-apiserver.yaml
```


### Create a personal access token in github and put your github personal access token to kubeconfig

```shell
vi ~/.kube/config
```
```yaml
- name: pengpeng404
  user:
    token: <mytoken>
```
```shell
k get po --user pengpeng404

#Error from server (Forbidden): pods is forbidden: User "pengpeng404" cannot list resource "pods" in API group "" in the namespace "default"

#root@master:/home/cadmin/authn-webhook# ./bin/amd64/authn-webhook
#2024/10/15 10:29:40 receving request
#2024/10/15 10:29:41 [Success] login as pengpeng404
```

### Reset the env

```sh
cp ~/kube-apiserver.yaml /etc/kubernetes/manifests/kube-apiserver.yaml
```



































