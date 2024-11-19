### 守护线程


DaemonSet 特点：

- 1、DaemonPod 运行在集群中的每个节点上
- 2、每个节点之后一个Pod实例
- 3、新的节点加入集群，该Pod自动创建，删除节点，相应的收回Pod
- 4、与Deployment不同，DaemonSet直接操作 Pod


DaemonSet demo

- 1、网络插件的 Agent 组件，用来处理这个节点上的容器网络
- 2、存储插件的 Agent 组件，用来在这个节点上挂载远程存储目录
- 3、各种监控组件和日志组件


DaemonSet 会为 Pod 增加调度相关的字段 tolerations nodeAffinity，

对于 nodeAffinity，
DaemonSet为 Pod 增加一个nodeAffinity，
保证这个Pod只会在指定的节点启动

对于tolerations，
当一个节点没有安装网络插件的时候，
会被自动打上 `node.kubernetes.io/network-unavailable` taint，
而DaemonSet会为 Pod 增加这个 toleration，
这样网络插件的组件就会调度到这个机器上运行起来


attention：

对DaemonSet进行回滚操作，
Revision不会回退，
而是增加

每一次的升级操作，
都会创建一个 `ControllerRevision` 对象，
用来保存每次升级的记录，
DaemonSet和StatefulSet都是这样









































