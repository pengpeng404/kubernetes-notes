
# 节点资源检测



## node-problem-detector

- Runtime 无响应


本身是个 daemonSet


```yaml
# npd-ds.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: node-problem-detector-v0.1
  labels:
    k8s-app: node-problem-detector
    version: v0.1
    kubernetes.io/cluster-service: "true"
spec:
  selector:
    matchLabels:
      k8s-app: node-problem-detector
      version: v0.1
      kubernetes.io/cluster-service: "true"
  template:
    metadata:
      labels:
        k8s-app: node-problem-detector
        version: v0.1
        kubernetes.io/cluster-service: "true"
    spec:
      hostNetwork: true
      containers:
        - name: node-problem-detector
          image: chainguard/node-problem-detector
          securityContext:
            privileged: true
          resources:
            limits:
              cpu: "200m"
              memory: "100Mi"
            requests:
              cpu: "20m"
              memory: "20Mi"
          volumeMounts:
            - name: log
              mountPath: /log
              readOnly: true
      volumes:
        - name: log
          hostPath:
            path: /var/log/
```


```shell
k create -f npd-ds.yaml
# 监听故障
sudo sh -c "echo 'kernel: BUG: unable to handle kernel NULL pointer dereference at TESTING' >> /dev/kmsg"

# 查看信息发现没有 因为 npd 需要权限更新节点信息
# 使用 Helm 安装试一下
```


## 节点问题排查手段


```shell
# 查看 logs 排查
kubectl logs -f <pod-name>
# 查看其中一个容器的日志
kubectl logs -f -c <container-name> <pod-name>
# 查看上一次的日志
kubectl logs -f <pod-name> --previous

# 针对 systemd 拉起的服务
journalctl -afu kubelet -S "2020-09-09 17:00:00"
```


















































