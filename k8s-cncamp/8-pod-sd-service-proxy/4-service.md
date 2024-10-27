# service 对象

```yaml
spec:
  clusterIP: 10.98.170.73
  clusterIPs:
  - 10.98.170.73
  externalTrafficPolicy: Cluster
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - nodePort: 30007
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx
  sessionAffinity: None
  type: NodePort
```

- selector 过滤查询 选择 pod 找到对应的 IP
- port 服务虚 IP 的端口
- targetPort 真实服务器端口

endPoint 是一个中间表 描述 pod 和 service 的映射关系

```yaml
# nginx-deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
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
          readinessProbe:
            exec:
              command:
                - cat
                - /tmp/healthy
            initialDelaySeconds: 5
            periodSeconds: 5
```

```yaml
# service.yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-basic
spec:
  type: ClusterIP
  ports:
    - port: 80
      protocol: TCP
      name: http
  selector:
    app: nginx

```

```shell
k get ep nginx-basic -oyaml
#subsets:
#- notReadyAddresses:
#  - ip: 10.244.166.135
#    nodeName: node1
#    targetRef:
#      kind: Pod
#      name: nginx-deployment-b7f55bbc9-vk5tt
#      namespace: default
#      uid: 226a7241-3bb7-4e1e-bb2b-440303641e47
#  ports:
#  - name: http
#    port: 80
#    protocol: TCP

# 只创建了一个 pod 并且是 notReady

k get po -owide
#NAME                               READY   STATUS    RESTARTS   AGE     IP               NODE    NOMINATED NODE   READINESS GATES
#nginx-deployment-b7f55bbc9-vk5tt   0/1     Running   0          4m39s   10.244.166.135   node1   <none>           <none>
```

此时 service 不可访问 因为后端没有 ready 的 pod

当创建一个 service 带有 selector，
那么 endPoint controller 创建一个同名的 endpoint 对象，
把所有的 pod 放到自己的地址栏中

```shell
# scale up to 2 and then
k exec nginx-deployment-b7f55bbc9-vk5tt -- touch /tmp/healthy
k exec nginx-deployment-b7f55bbc9-r9mzd -- touch /tmp/healthy
k get ep
#NAME          ENDPOINTS                          AGE
#kubernetes    172.30.219.96:6443                 161m
#nginx-basic   10.244.166.135:80,10.244.76.3:80   10m

k get po -owide
#NAME                               READY   STATUS    RESTARTS   AGE   IP               NODE    NOMINATED NODE   READINESS GATES
#nginx-deployment-b7f55bbc9-r9mzd   1/1     Running   0          56s   10.244.76.3      main    <none>           <none>
#nginx-deployment-b7f55bbc9-vk5tt   1/1     Running   0          11m   10.244.166.135   node1   <none>           <none>

# 此时 endpoint 显示已经 ready 的 pod ip
# 这些 ip 可以接受流量
# notReady 的 pod 是不会接受流量的
```


### 不定义 Selector 的 Service

- EndPoint Controller 不会自动创建 EndPoint
- 用户可以手动创建 EndPoint 设置任意 IP



## Service 类型

```shell
cat /etc/kubernetes/manifests/kube-apiserver.yaml
#- --service-cluster-ip-range=10.96.0.0/12
# 由 apiserver 分配
```

- ClusterIP
- nodePort
- LoadBalancer
- Headless Service
- ExternalName

### Service Topology




































