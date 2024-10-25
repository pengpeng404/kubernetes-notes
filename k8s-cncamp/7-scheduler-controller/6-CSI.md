# CSI

如何写文件？


## 临时存储

### emptyDir

就是一个空的卷 和 pod 生命周期紧紧绑定

```yaml
# nginx.yaml
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
          image: nginx:1.27.0
          volumeMounts:
          - mountPath: /cache
            name: cache-volume
      volumes:
      - name: cache-volume
        emptyDir: {}
```

```shell
k create -f nginx.yaml
crictl ps |grep ng-744478f867-hjblm
crictl inspect ecddec44ac826
#{
#    "destination": "/cache",
#    "type": "bind",
#    "source": "/var/lib/kubelet/pods/975913c1-2353-4c06-9932-5faf7a396cfc/volumes/kubernetes.io~empty-dir/cache-volume",
#    "options": [
#      "rbind",
#      "rprivate",
#      "rw"
#    ]
#}

#这个不是 overlayFS 和写在主机上的效率是一样的
#其次 当删除 pod kubelet会清除这个目录
# 随着pod生命周期来的

```


## 半持久化存储

### hostPath

可以在容器内部操作主机上的文件

主机上的目录文件和pod生命周期是解耦的

如果不清理 主机上的文件还在

```shell





```

如果 pod 发生飘逸 这个数据就找不到了

如果删除 pod 那么这个数据就在主机上 如果没人清理 就一直在

不能把 hostPath 的权限开放出来 很多人会动主机的数据




### pv pvc




## Rook



































