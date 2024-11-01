
# kind 安装集群


## kind 配置

```shell
kind create cluster --name new-cluster --config kind-config.yaml
```

```yaml
# kind-config.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    image: kindest/node:v1.31.1
  - role: worker
    image: kindest/node:v1.31.1
  - role: worker
    image: kindest/node:v1.31.1
```

## kubectl 配置

```shell
# 查看当前使用的上下文
kubectl config current-context

# 查看所有可用上下文
kubectl config get-contexts

# 切换到指定 kind 集群
kubectl config use-context kind-my-cluster

# 验证链接的集群
kubectl cluster-info
```























































