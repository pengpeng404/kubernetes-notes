# etcd start


`etcd` 启动参数
- `name` 
  - default
- `data-dir`
  - default.etcd/member/wal
  - default.etcd/member/snap
  - snapshot 给 etcd 做镜像
- `listen-peer-urls`
- `listen-client-urls`


## 容量管理

单个对象不建议超过 `1.5M`

默认容量 `2G`

不建议超过 `8G`

其中 `boltdb` 会把整个数据库信息映射到内存里 如果容量是 `8G` 内存已经占用 `8G`

还有 `KVIndex` 如果 `boltdb` 很大 对于内存的开销很大
```log
5000 Node 200000 Pods 3000 Service 30000 Namespace
共占用 800MB
```


## Alarm

```shell
# 设置 etcd 存储大小
etcd --quota-backend-bytes=$((16*1024*1024))

# 写爆磁盘
while [ 1 ]; do dd if=/dev/urandom bs=1024 count=1024 | ETCDCTL_API=3 etcdctl put key || break; done

# 查看 endpoint 状态
# 如果磁盘爆了 变成只读模式
ETCDCTL_API=3 etcdctl --write-out=table endpoint status

# 查看 alarm
ETCDCTL_API=3 etcdctl alarm list

# 空间压缩 清理碎片 做碎片整理
## 有自动 compact 的启动参数
ETCDCTL_API=3 etcdctl compact 3
ETCDCTL_API=3 etcdctl defrag

# 清理 alarm
ETCDCTL_API=3 etcdctl alarm disarm
```


















