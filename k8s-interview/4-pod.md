### pod 对象

重启pod V.S. 重建pod

重启pod

pod中的容器重新启动，
而pod对象本身不会被删除或者重建

- 1、容器 CrashLoop 根据 restartPolicy 重新启动容器
- 2、LivenessProbe
- 3、节点资不足 如内存不足 触发 OOM
- 4、


重建pod

pod被删除，
创建一个新的pod，
pod的UID发生变化

- 1、pod所属的控制器滚动更新
- 2、pod被驱逐 比如节点资源不足触发pod驱逐策略
- 3、pod配置更新 比如 deployment 模板更新
- 4、抢占 高优先级的pod需要调度 当前节点没有足够资源 低优先级的pod会被抢占和删除


如何理解：

- 1、重启pod是pod内的容器出问题 进程出问题 重启逻辑由 kubelet 执行
- 2、重建pod是控制器决定，是对pod对象的完全替换




































