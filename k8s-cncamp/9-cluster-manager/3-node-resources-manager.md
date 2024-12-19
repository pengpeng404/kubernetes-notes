# 节点资源管理

## 状态上报

kubelet周期性向apiServer汇报，
并更新节点的相关健康和资源使用信息

- 节点基础信息
  - IP 地址
  - 操作系统
  - 内核
  - 运行时
  - kubelet版本信息
  - kube-proxy版本信息
- 节点资源信息
  - CPU
  - 内存
  - HugePage
  - 临时存储
  - GPU等

调度器在为Pod选择节点的时候会将机器的状态信息作为依据

#### Lease

在早期版本kubelet状态上报直接更新node对象，
上报的信息包含状态信息和资源信息，
因此数据包很大，
给apiserver和etcd造成了很大的压力。

后来引入Lease对象用来保存健康信息，
在默认的40s的`nodeLeaseDurationSeconds`周期内，
若Lease对象没有被更新，
则对应节点可以被判定为不健康。

也就是说，
把节点资源信息和节点健康信息分离，
使用Lease来上报健康信息。




```shell
k get lease node1 -oyaml -n kube-node-lease
#apiVersion: coordination.k8s.io/v1
#kind: Lease
#metadata:
#  creationTimestamp: "2024-10-26T03:38:21Z"
#  name: node1
#  namespace: kube-node-lease
#  ownerReferences:
#  - apiVersion: v1
#    kind: Node
#    name: node1
#    uid: 5ecaaec7-3688-4c69-8b69-1822ff352fdc
#  resourceVersion: "395334"
#  uid: 7a373a2f-a3a6-44b3-b8f1-fdd957da684f
#spec:
#  holderIdentity: node1
#  leaseDurationSeconds: 40
#  renewTime: "2024-10-29T09:29:40.840354Z"

## 你会看到 Lease 对象 name 和 node 名字一致
## Lease 存在在 kube-node-lease ns 中
```

每一个node都有一个Lease对象，
Lease对象就是用来保持心跳的对象，
但是他不携带任何上报信息，
只用来保持心跳，
如果信息没有及时更新，
则认为这个node不健康

kubelet通过PATCH或者PUT请求更新非心跳信息，
如节点资源使用或者节点状态变化，
不会与Lease对象混用

在心跳的过程中，
kubelet仅会对Lease对象做一次PATCH请求，
用于更新心跳相关字段

- holderIdentity 谁拿到这个 lease
- leaseDurationSeconds 这个 lease 持有多久
- renewTime 上次更新的时间

如果超过这个时间 kubelet 还没有 renew 则说明节点不健康



## 资源预留


- Capacity
  - 节点资源能力
  - `cat /proc/cpuinfo`
  - `cat /proc/meminfo`
  - `kube-reserved` + `system-reserved` + `eviction-threshold`
- Allocatable



节点磁盘管理

- nodefs
  - 工作目录和容器日志
  - `ls /var/lib/kubelet/pods/`
- imagefs
  - `ls /var/lib/containerd/`



## 驱逐管理

不可压缩资源 kubelet 自保 做驱逐管理 终止一些容器进程



## 资源可用额监控


![qos](images/qos.png)


## OOM Killer

```yaml



```











































