# apiserver 高可用


### apiserver 多副本

`apiserver` 是无状态的 `restServer`

无状态 所以方便 `scale up/down`

负载均衡：
- 在多个 `apiserver` 实例之上 配置负载均衡 提供一个唯一入口
- 证书可能需要加上 `loadBalancer VIP` 重新生成


推荐使用静态 `pod` 来管理 `apiserver`，
因为本身 `kubelet` 就是保证应用高可用的，
并且 `kubelet` 可以做探活 比 `systemd` 功能多


### 预留充足的 cpu 内存资源

规划 `apiserver` 资源时 要为未来预留充分资源

### 限流一定要管

APF 是默认开启的

先按照集群规模设置相应的 `max-request-inflight` `max-mutating-request-inflight`


### 设置合适的缓存大小

`apiserver` 对 `etcd` 有保护

`apiserver` 有 `etcd` 的 `watch` 缓存

设置适当大小的 `watch-cache-sizes`



### 监听大于轮询

`list` 的过滤是在 `apiserver` 中进行的 频繁的轮询会压死 `apiserver`


### 如何访问 apiserver

对于外部用户，永远只通过 `loadBalancer` 访问

`apiserver` 提供 对内的访问入口 和 对外的访问入口

我们的 `kubernetes` 组件最好使用统一的访问入口

因为所有的控制器都在监听 如果有一个配置错误 可能整个集群都掉了


## 搭建多租户的 Kubernetes 集群

- 授信
  - 认证
    - 禁止匿名访问 只允许可信用户操作
    - `serviceAccount`
    - 外部认证平台 `webhook`
  - 授权
    - 防止多用户之间互相影响
    - 普通用户 管理员
    - 自动化方案 建立 `namespace` 的时候自动控制器绑定权限
    - 自动成为这个 `namespace` 的 `admin`
- 隔离
  - 可见性隔离
    - 用户只关心自己的应用 无需看到其他用户的服务和部署
  - 资源隔离
    - 有些关键项目对资源要求较高 需要专有设备 不与他人共享
    - 此时可以使用 `taint` 给专用节点 `node` 打上污点 这个服务需要容忍这个 `taint`
    - 其他人不知道这个节点有这个 `taint` 就不会调度到这个节点上去
  - 应用访问隔离
    - 部署的服务不想让其他人访问
- 资源管理
  - `Quota` 管理
    - 使用 `resourceQuota` 来限制每个 `namespace` 里能够使用多少资源
    - 可以写 `resourceController` 当用户建好 `namespace` 自动把 `resourceQuota` 建出来
    - 默认限制对象 自动化


























































