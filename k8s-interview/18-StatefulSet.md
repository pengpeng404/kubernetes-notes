### StatefulSet


如何理解 StatefulSet

1、拓扑状态
- 多个应用实例不是完全对等的关系
- 严格按照某些顺序启动 删除 重建 也是按照顺序
- 重建之后网络标识必须和原来一样

2、存储状态
- 每个实例是绑定不同的存储数据的

#### 拓扑状态

Service 被访问的方式：
- 1、Service VIP （Virtual IP）
- 2、Service DNS my-svc.my-namespace-svc.cluster.local


而在第二种DNS方式下，还可以分为两种处理方式：
- Normal Service
  - 访问my-svc.my-namespace-svc.cluster.local解析到的是Service 的VIP 就和第一种一样了
- Headless Service
  - 访问my-svc.my-namespace-svc.cluster.local解析到的是代理的某一个pod的ip地址 直接解析到具体的pod ip


可以通过 Headless Service 来为pod 分配唯一的可解析身份

`<pod-name>.<service-name>.<namespace>.svc.cluster.local`


StatefulSet在创建的时候，
按照编号一个一个启动，
只有之前的 Ready 之后，
才会创建第二个，
不然第二个一直 pending 状态，
严格按照拓扑结构启动，
所以设置合理的探针非常重要


如果之前的pod出现错误，
需要重新启动，
那么后面的pod会被杀掉，
等待之前的pod恢复之后，
才会启动

同时需要通过 Headless Service 来访问 StatefulSet 应用，
使用 DNS 来提供一个同一的访问入口


#### 存储状态




StatefulSet总结：

1、StatefulSet的控制器直接管理的是 Pod，
因为每个pod的hostname、名字都是不同的，
携带编号

2、Kubernetes通过HeadlessService，
为这些有编号的 Pod在DNS服务器中生成带有同样编号的DNS记录，
这条记录解析出来的pod IP 地址，
会随着后端的pod删除和重建而自动更新，
这是 Service机制本身的能力

3、StatefulSet会为每一个pod分配并创建一个同样的编号的PVC，
保证每个pod都拥有一个独立的Volume


#### StatefulSet 把 MySQL 容器化


主从复制的 MySQL 集群
- 1、一个主节点 多个从节点
- 2、从节点需要能水平扩展
- 3、所有的写操作 只能在主节点上执行
- 4、读操作可以在所有节点上执行











