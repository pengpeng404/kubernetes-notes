# 扩展资源

扩展资源是 kubernetes.io 域名之外的标准资源名称

自定义扩展资源无法使用 kubernetes.io 作为资源域名




## 管理扩展资源



```shell
cat ~/.kube/config
echo <key/cert> |base64 -d > <admin.crt/admin.key>

curl --key admin.key --cert admin.crt --header "Content-Type: application/json-patch+json" \
     --request PATCH -k \
     --data '[{"op": "add", "path": "/status/capacity/cncamp.com~1reclaimed-cpu", "value": "2"}]' \
     https://172.30.219.96:6443/api/v1/nodes/node1/status
     
k get no -oyaml
#  status:
#    addresses:
#    - address: 172.30.219.97
#      type: InternalIP
#    - address: node1
#      type: Hostname
#    allocatable:
#      cncamp.com/reclaimed-cpu: "2"
#      cpu: "2"
#      ephemeral-storage: "37694649077"
#      hugepages-1Gi: "0"
#      hugepages-2Mi: "0"
#      memory: 1610772Ki
#      pods: "110"
#    capacity:
#      cncamp.com/reclaimed-cpu: "2"
#      cpu: "2"
#      ephemeral-storage: 40901312Ki
#      hugepages-1Gi: "0"
#      hugepages-2Mi: "0"
#      memory: 1713172Ki
#      pods: "110"
```

```yaml
# nginx-reclaimed.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-reclaimed
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx
          resources:
            limits:
              cncamp.com/reclaimed-cpu: 3
            requests:
              cncamp.com/reclaimed-cpu: 3
```

```shell
k create -f nginx-reclaimed.yaml
k describe po nginx-reclaimed-6d5db85d8b-7kb4n
#vents:
#  Type     Reason            Age   From               Message
#  ----     ------            ----  ----               -------
#  Warning  FailedScheduling  15s   default-scheduler  0/2 nodes are available: 2 Insufficient cncamp.com/reclaimed-cpu. preemption: 0/2 nodes are available: 2 No preemption victims found for incoming pod.

# edit to 1
#nginx-reclaimed-cfccfc498-62zq9    1/1     Running   0              64s


```

edit 只能修改 spec 不能修改 status status 只能 patch


## 可扩展资源应用点 降本增效

一般的应用有波峰波谷，
然而用户在使用的时候大多数把 pod 资源完全定义为 Guaranteed，
这样做 cpu 利用率很低，
可以把 cpu 资源监控为一个可自定义资源，
当应用处于波谷的时候，
即使应用是 Guaranteed，
也可以在该 node 上增加自定义资源，
把其他应用调度到这个 node，
增加 cpu 利用率，
离开波谷的时间段，
自定义资源缩容即可


























































