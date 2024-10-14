
### Install helm

https://github.com/helm/helm/releases
```shell
tar -zxvf helm-v3.0.0-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm
helm help
```

### Download bitnami etcd

```sh
helm repo add bitnami https://charts.bitnami.com/bitnami
helm pull bitnami/etcd
tar -xvf etcd-6.8.4.tgz
vi etcd/values.yaml
```

And set persistence to false:

```yaml
persistence:
  ## @param persistence.enabled If true, use a Persistent Volume Claim. If false, use emptyDir.
  ##
  enabled: false
```

### Install etcd by helm chart

```sh
helm install my-release ./etcd
```

### Start etcd client

```sh
kubectl run my-release-etcd-client --restart='Never' --image docker.io/bitnami/etcd:3.5.16-debian-12-r1 --env ROOT_PASSWORD=$(kubectl get secret --namespace default my-release-etcd -o jsonpath="{.data.etcd-root-password}" | base64 -d) --env ETCDCTL_ENDPOINTS="my-release-etcd.default.svc.cluster.local:2379" --namespace default --command -- sleep infinity
```

```sh
kubectl exec --namespace default -it my-release-etcd-client -- bash
etcdctl --user root:$ROOT_PASSWORD put /message Hello
etcdctl --user root:$ROOT_PASSWORD get /message
```























