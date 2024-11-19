### watch

从这里开始记录源码的阅读笔记，
整个kubernetes源码非常复杂，
但是之前有了整体的概念认知，
阅读代码会轻松很多，
加深对与kubernetes的具体的理解，
从抽象的概念，
到具体的实现


1. API Server：事件生成与分发
   源码模块：

事件生成：

目录：staging/src/k8s.io/apiserver/pkg/storage
文件：etcd3/store.go
关键函数：

Watch()：核心的 Watch 逻辑，负责监听资源变化并生成事件。
事件分发：

目录：staging/src/k8s.io/apiserver/pkg/storage/cacher
文件：cacher.go
关键函数：

dispatchEvent()：将事件从 watchCache 中推送到监听者。
推荐阅读顺序：

etcd3/store.go 的 Watch() 实现，理解资源变更如何触发 Watch。
cacher.go 的 dispatchEvent()，理解事件如何通过缓存层被分发到 Watch 客户端。



2. Controller Manager：Informer 获取事件
   源码模块：

Informer 工厂：

目录：staging/src/k8s.io/client-go/informers
文件：factory.go
关键函数：

Start()：启动 Informer 的逻辑。
SharedInformerFactory：创建和管理共享 Informer。
Informer 核心：

目录：staging/src/k8s.io/client-go/tools/cache
文件：
shared_informer.go：SharedInformer 的实现。
reflector.go：负责从 API Server 监听事件。
delta_fifo.go：负责事件的本地缓存和队列。
关键函数：

NewSharedInformer()：创建一个新的 SharedInformer。
Run()：启动事件监听和缓存。
ListAndWatch()：通过 list 获取全量数据并 watch 接收增量事件。
推荐阅读顺序：

从 factory.go 的 Start() 方法入手，了解 SharedInformerFactory 如何启动。
阅读 shared_informer.go 的 Run() 方法，理解 Informer 的整体运行机制。
深入 reflector.go 的 ListAndWatch() 方法，理解事件监听的细节。



3. Controller：事件处理与业务逻辑
   源码模块：

Deployment Controller：

目录：pkg/controller/deployment
文件：deployment_controller.go
关键函数：

addDeployment()：处理新增 Deployment 的事件。
updateDeployment()：处理 Deployment 的更新事件。
deleteDeployment()：处理 Deployment 的删除事件。
WorkQueue：

目录：staging/src/k8s.io/client-go/util/workqueue
文件：workqueue.go
关键函数：

Add()：将事件加入队列。
Get()：从队列中获取事件。
Done()：标记事件处理完成。
推荐阅读顺序：

从 deployment_controller.go 的 NewDeploymentController() 方法入手，了解 Deployment Controller 的初始化。
阅读 addDeployment() 等方法，理解事件如何从 Informer 传递到业务逻辑。
深入 WorkQueue 的实现，理解事件的分发和调度机制。



apiserver中一个store对象，
该对象是对集群中的所有对象的一个抽象概念，
无论是ns的还是非ns的，
在apiserver运行的过程中，
都会对不同的对象创建不同的store对象，
但是这个对象是一种对象，
不是这对于单个具体的对象，
比如 pod、node、deployment、service等

```go
// staging/src/k8s.io/apiserver/pkg/storage/etcd3/store.go
type store struct {
	client              *kubernetes.Client
	codec               runtime.Codec
	versioner           storage.Versioner
	transformer         value.Transformer
	pathPrefix          string
	groupResource       schema.GroupResource
	groupResourceString string
	watcher             *watcher
	leaseManager        *leaseManager
	decoder             Decoder
}

// /registry/pods/
// /registry/nodes/
```

那么apiserver在启动的过程中，
存储的初始化，
也就是store初始化是通过StorageFactory完成的，
她负责为每种资源类型创建store对象，




































