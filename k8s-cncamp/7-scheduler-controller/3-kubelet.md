# kubelet

![kubelet](images/kubelet.png)


![kubeletpod](images/kubeletpod.png)

每个节点上都运行一个 kubelet 服务进程， 默认监听 10250 端口
- 接受并执行 master 发来的指令
- 管理 pod 以及 pod 中的容器
- 每个 kubelet 进程会在 apiserver 上注册节点自身信息 定期向 master 节点汇报资源使用情况
- 并通过 cAdvisor 监控节点和容器资源




kubelet 本身也是一个 controller 模式

关注 pod update 事件

/file 静态加载pod

apiserver 


为什么kubernetes要给每个node设置pod上限

为什么不可以跑很多的pod

答案是不可以

因为需要 relist 这个动作是有开销的 不断发 list 请求给 运行时

量级到了一定程度 变得复杂


### 节点管理


### Pod 管理


### Pod 启动流程



为什么要运行一个 sandbox

所谓的 pod 是一组容器的组合 这组容器默认共享 ns 共享系统层面资源

引入 sandbox 先启动一个 sandbox  pause

这个 pause 是个永远 sleep 的 不消耗 cpu资源 并且镜像很小 极度稳定

网络的配置可以配置在这里 并且这个容器不会退出

作为整个pod的底座 把网路挂上去

还有 init 容器 也需要网络配置 就是说在容器启动之前 就已经需要有网络配置条件了

必须有网络就绪的状态 先把网络挂载

所有的网络都是配置在 sandbox 上的




- checkAdmit
  - 可以在kubelet配置准入插件 当一定的条件满足之后再启动这个pod
- 检查网络插件
- 等待所有的 volume 就绪 比如外挂存储 configmap 等 cni cri csi 谁先启动 csi
- log


































































