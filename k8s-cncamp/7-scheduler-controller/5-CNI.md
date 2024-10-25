# CNI

- pod - pod pod- node --> CNI
- pod - service - pod --> kube-proxy
- 外部流量入站 --> ingress


kubernetes 网络模型设计的基础原则是：
- 所有的 pod 能够不通过 NAT 就能互相访问
- 所有的 node 能够不通过 NAT 就能互相访问
- 容器内看到的 IP 地址和外部组件看到的容器 IP 是一样的

IP 地址是以 pod 为单位进行分配的 每个 pod 都有一个独立的 IP 地址 Pod 内的所有容器
可以通过 localhost:port 访问对方



### CNI 插件分类和常见插件
- IPAM：IP 分配
- 主插件：网卡配置
  - bridge




### CNI 插件运行机制

```shell
# 默认的 CNI 配置目录
cd /etc/cni/net.d/
cat 

# 编译好的二进制文件
cd /opt/cni/bin/

# kubelet 掉用 cni 接口的时候 调用了本地的二进制文件
# 


```

比如要启动一个容器，运行时要调用 cni 插件 调用的时候是个链

先调用 ipam 做 ip 分配 为这个 pod 分配一个 ip

主的plugin 会把这个 ip 分配给这个 容器的 ns

然后把结果告诉 cri 

containerruntime就会把这个ip告诉 kubelet 再上报给apiserver

apiserver把ip写入状态





### 打通主机网络




### CNI Plugin

- Calico
- Cilium 看好 规模不是很大 效率更好




## Calico

```shell
ks get ds
#NAME          DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
#calico-node   3         3         3       3            3           kubernetes.io/os=linux   32d
#kube-proxy    3         3         3       3            3           kubernetes.io/os=linux   32d

ks get po calico-node-cznt4 -oyaml
#initContainers:
#- command:
#    - /opt/cni/bin/install
#    env:
#    - name: CNI_CONF_NAME
#      value: 10-calico.conflist
#    - name: CNI_NETWORK_CONFIG
#      valueFrom:
#        configMapKeyRef:
#          key: cni_network_config
#          name: calico-config

# 负责把 Calico 的插件复制到主机 cni 目录
# 使用一个 init container 把文件拷贝到主机文件目录




```




### Calico VXLan

主机上有一个 vxlan 设备

```shell
ip a


```


一步一步看数据面

```shell
k get crd
#ipamblocks.crd.projectcalico.org                      2024-09-21T03:18:58Z
#ipamconfigs.crd.projectcalico.org                     2024-09-21T03:18:58Z
#ipamhandles.crd.projectcalico.org                     2024-09-21T03:18:58Z
#ippools.crd.projectcalico.org                         2024-09-21T03:18:58Z

k get ippools.crd.projectcalico.org -oyaml
#  spec:
#    allowedUses:
#    - Workload
#    - Tunnel
#    blockSize: 26
#    cidr: 10.244.0.0/16
#    ipipMode: Always
#    natOutgoing: true
#    nodeSelector: all()
#    vxlanMode: Never

# 默认开启 ipip 模式 vxlan 关闭
# blockSize: 26 在每个主机上会分多大的 ip 段

k get ipamblocks.crd.projectcalico.org -oyaml
# 用来记录每个节点的 cidr 以及分出去的 ip
# 同时也记录了该 ip 是哪个节点上的哪个 pod

k get ipamhandles.crd.projectcalico.org
#NAME                                                                               AGE
#ipip-tunnel-addr-master                                                            32d
#ipip-tunnel-addr-node1                                                             32d
#ipip-tunnel-addr-node2                                                             32d
#k8s-pod-network.204a152f2a239238ef67ab9f043d7b7afffa49ef0a5a9ee3c32c87e860f5bc71   9h
#k8s-pod-network.6cb1a7a3e006e335302e9865073e0a72c3df9f49e659a5032aaf9c6326651f3c   9h
#k8s-pod-network.ae1bbab2fc2b19fd0ddb1b6a91ad9db028df62ff82b7e77c004aca2849eecc81   9h
#k8s-pod-network.b5c5999917550d81f99f3bcbc54f8d644f2790ed6152c1427336b4a803dedd5e   9h
#k8s-pod-network.c9f37ce1d27601d9a85d45bc222e0bcde939f96a77fe20c8086d963e6fa249b0   9h

k get ipamhandles.crd.projectcalico.org k8s-pod-network.204a152f2a239238ef67ab9f043d7b7afffa49ef0a5a9ee3c32c87e860f5bc71 -oyaml
#apiVersion: crd.projectcalico.org/v1
#kind: IPAMHandle
#metadata:
#  annotations:
#    projectcalico.org/metadata: '{"creationTimestamp":null}'
#  creationTimestamp: "2024-10-23T01:09:58Z"
#  generation: 1
#  name: k8s-pod-network.204a152f2a239238ef67ab9f043d7b7afffa49ef0a5a9ee3c32c87e860f5bc71
#  resourceVersion: "718586"
#  uid: 7b898ad0-89a0-4cbe-a124-717e9cce2b1a
#spec:
#  block:
#    10.244.219.64/26: 1
#  deleted: false
#  handleID: k8s-pod-network.204a152f2a239238ef67ab9f043d7b7afffa49ef0a5a9ee3c32c87e860f5bc71
```

整个数据链路是怎么通的

BGP 互相交换路由

```shell
kubectl run -it net --image=docker.io/library/nicolaka/netshoot:v0.1 -- /bin/bash
k get po -owide
#NAME                       READY   STATUS             RESTARTS      AGE   IP               NODE    NOMINATED NODE   READINESS GATES
#apphttp-5887bfd99d-qpctj   1/1     Running            1 (23m ago)   10h   10.244.104.63    node2   <none>           <none>
#net                        1/1     Running            1 (7s ago)    9s    10.244.104.7     node2   <none>           <none>

k exec -it net -- /bin/bash
ping 10.244.104.63
#PING 10.244.104.38 (10.244.104.38) 56(84) bytes of data.
#64 bytes from 10.244.104.38: icmp_seq=1 ttl=63 time=0.056 ms
#64 bytes from 10.244.104.38: icmp_seq=2 ttl=63 time=0.053 ms
#64 bytes from 10.244.104.38: icmp_seq=3 ttl=63 time=0.047 ms

ip a
#1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
#    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
#    inet 127.0.0.1/8 scope host lo
#       valid_lft forever preferred_lft forever
#    inet6 ::1/128 scope host 
#       valid_lft forever preferred_lft forever
#2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
#    link/ipip 0.0.0.0 brd 0.0.0.0
#4: eth0@if14: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1480 qdisc noqueue state UP group default 
#    link/ether 92:93:8a:d9:90:2d brd ff:ff:ff:ff:ff:ff link-netnsid 0
#    inet 10.244.104.7/32 scope global eth0
#       valid_lft forever preferred_lft forever
#    inet6 fe80::9093:8aff:fed9:902d/64 scope link 
#       valid_lft forever preferred_lft forever

ip r
#default via 169.254.1.1 dev eth0 
#169.254.1.1 dev eth0 scope link 

arping 169.254.1.1
#ARPING 169.254.1.1 from 10.244.104.7 eth0
#Unicast reply from 169.254.1.1 [EE:EE:EE:EE:EE:EE]  0.893ms
#Unicast reply from 169.254.1.1 [EE:EE:EE:EE:EE:EE]  0.521ms

# 从 eth0 出去 到达 veth 的另外的一个口

ip a
# 主机
#8: calif66aa2997ff@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1480 qdisc noqueue state UP group default 
#    link/ether ee:ee:ee:ee:ee:ee brd ff:ff:ff:ff:ff:ff link-netns cni-3259448b-763b-90d8-b24f-e8a0e3770dea
#14: cali8eecb1f59c6@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1480 qdisc noqueue state UP group default 
#    link/ether ee:ee:ee:ee:ee:ee brd ff:ff:ff:ff:ff:ff link-netns cni-ce8c08d9-2e65-3e46-4bdc-97f015f55ab8

ip r
#default via 10.0.2.1 dev enp0s3 proto dhcp src 10.0.2.12 metric 100 
#10.0.2.0/24 dev enp0s3 proto kernel scope link src 10.0.2.12 metric 100 
#10.0.2.1 dev enp0s3 proto dhcp scope link src 10.0.2.12 metric 100 
#blackhole 10.244.104.0/26 proto bird 
#10.244.104.7 dev cali8eecb1f59c6 scope link 
#10.244.104.63 dev calif66aa2997ff scope link 
#10.244.166.128/26 via 192.168.34.102 dev tunl0 proto bird onlink 
#10.244.219.64/26 via 192.168.34.101 dev tunl0 proto bird onlink 
#192.168.34.0/24 dev enp0s8 proto kernel scope link src 192.168.34.103 

# 这就是同主机下不同 pod 之间如何通讯 在主机上的路由表中直接传到 veth 的口
# 直接到目标 pod


```



```shell
# 跨主机 pod 如何联通

# 每个节点都会跑一个 bird 的 daemon
# 不同的计算节点 bird 有长链接 互相交换彼此的路由表
# 同样的

ip r
#default via 10.0.2.1 dev enp0s3 proto dhcp src 10.0.2.12 metric 100 
#10.0.2.0/24 dev enp0s3 proto kernel scope link src 10.0.2.12 metric 100 
#10.0.2.1 dev enp0s3 proto dhcp scope link src 10.0.2.12 metric 100 
#blackhole 10.244.104.0/26 proto bird 
#10.244.104.7 dev cali8eecb1f59c6 scope link 
#10.244.104.63 dev calif66aa2997ff scope link 
#10.244.166.128/26 via 192.168.34.102 dev tunl0 proto bird onlink 
#10.244.219.64/26 via 192.168.34.101 dev tunl0 proto bird onlink 
#192.168.34.0/24 dev enp0s8 proto kernel scope link src 192.168.34.103 

#10.244.166.128/26 via 192.168.34.102 dev tunl0 proto bird onlink 
#10.244.219.64/26 via 192.168.34.101 dev tunl0 proto bird onlink 

# 这两条数据就是通过 bird daemon 互相交换路由表来的
# blackhole 10.244.104.0/26 proto bird 
# 读这个 balckhole 用来记录这个主机上是有这个子网网段的

```


### CNI plugin 对比













































