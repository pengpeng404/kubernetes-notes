# 高可用 etcd 集群搭建

高可用 `Kubernetes` 首先必须实现高可用 `etcd`

高可用 `etcd`
- 1、多个 `etcd member` 形成集群 保证数据安全性
- 2、做及时数据备份和故障恢复


### Install cfssl

```sh
apt install golang-cfssl
```

### Generate tls certs

```sh
mkdir -p /root/go/src/github.com/etcd-io

cd /root/go/src/github.com/etcd-io
```

```sh
git clone https://github.com/etcd-io/etcd.git
cd etcd/hack/tls-setup
```

### Edit req-csr.json and keep 127.0.0.1 and localhost only for single cluster setup.

```sh
vi config/req-csr.json
```

### Generate certs

```sh
apt install make
export infra0=127.0.0.1
export infra1=127.0.0.1
export infra2=127.0.0.1
make
mkdir /tmp/etcd-certs
mv certs /tmp/etcd-certs
```

### Start etcd cluster member

```sh
nohup etcd --name infra0 \
--data-dir=/tmp/etcd/infra0 \
--listen-peer-urls https://127.0.0.1:3380 \
--initial-advertise-peer-urls https://127.0.0.1:3380 \
--listen-client-urls https://127.0.0.1:3379 \
--advertise-client-urls https://127.0.0.1:3379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
--initial-cluster-state new \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra0.log &

nohup etcd --name infra1 \
--data-dir=/tmp/etcd/infra1 \
--listen-peer-urls https://127.0.0.1:4380 \
--initial-advertise-peer-urls https://127.0.0.1:4380 \
--listen-client-urls https://127.0.0.1:4379 \
--advertise-client-urls https://127.0.0.1:4379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
--initial-cluster-state new \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra1.log &

nohup etcd --name infra2 \
--data-dir=/tmp/etcd/infra2 \
--listen-peer-urls https://127.0.0.1:5380 \
--initial-advertise-peer-urls https://127.0.0.1:5380 \
--listen-client-urls https://127.0.0.1:5379 \
--advertise-client-urls https://127.0.0.1:5379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
--initial-cluster-state new \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra2.log &

```
### Member list

```sh
etcdctl --endpoints https://127.0.0.1:3379 \
  --cert /tmp/etcd-certs/certs/127.0.0.1.pem \
  --key /tmp/etcd-certs/certs/127.0.0.1-key.pem \
  --cacert /tmp/etcd-certs/certs/ca.pem \
  member list
```

### Put data

```sh
etcdctl --endpoints https://127.0.0.1:3379 \
  --cert /tmp/etcd-certs/certs/127.0.0.1.pem \
  --key /tmp/etcd-certs/certs/127.0.0.1-key.pem \
  --cacert /tmp/etcd-certs/certs/ca.pem \
  put a v1
  
etcdctl --endpoints https://127.0.0.1:3379 \
  --cert /tmp/etcd-certs/certs/127.0.0.1.pem \
  --key /tmp/etcd-certs/certs/127.0.0.1-key.pem \
  --cacert /tmp/etcd-certs/certs/ca.pem \
  get a -wjson
```


### Backup

```shell
# 在当前目录文件地址生成 snapshot.db 文件
etcdctl --endpoints https://127.0.0.1:3379 \
  --cert /tmp/etcd-certs/certs/127.0.0.1.pem \
  --key /tmp/etcd-certs/certs/127.0.0.1-key.pem \
  --cacert /tmp/etcd-certs/certs/ca.pem \
  snapshot save snapshot.db
```


### Kill process

```sh
ps -ef|grep "/tmp/etcd/infra"|grep -v grep|awk '{print $2}'|xargs kill
```

### Delete data

```sh
rm -rf /tmp/etcd
```

### Restore

```sh
# 把数据恢复 但是不启动 etcd
export ETCDCTL_API=3
etcdctl snapshot restore snapshot.db \
  --name infra0 \
  --data-dir=/tmp/etcd/infra0 \
  --initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-advertise-peer-urls https://127.0.0.1:3380

etcdctl snapshot restore snapshot.db \
    --name infra1 \
    --data-dir=/tmp/etcd/infra1 \
    --initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
    --initial-cluster-token etcd-cluster-1 \
    --initial-advertise-peer-urls https://127.0.0.1:4380

etcdctl snapshot restore snapshot.db \
  --name infra2 \
  --data-dir=/tmp/etcd/infra2 \
  --initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-advertise-peer-urls https://127.0.0.1:5380
```


### Restart

```shell
# 恢复的时候已经指定集群的信息 在重启的时候就不需要指定了
nohup etcd --name infra0 \
--data-dir=/tmp/etcd/infra0 \
--listen-peer-urls https://127.0.0.1:3380 \
--listen-client-urls https://127.0.0.1:3379 \
--advertise-client-urls https://127.0.0.1:3379 \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra0.log &

nohup etcd --name infra1 \
--data-dir=/tmp/etcd/infra1 \
--listen-peer-urls https://127.0.0.1:4380 \
--listen-client-urls https://127.0.0.1:4379 \
--advertise-client-urls https://127.0.0.1:4379 \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra1.log &

nohup etcd --name infra2 \
--data-dir=/tmp/etcd/infra2 \
--listen-peer-urls https://127.0.0.1:5380 \
--listen-client-urls https://127.0.0.1:5379 \
--advertise-client-urls https://127.0.0.1:5379 \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra2.log &

```


### Get data

```sh 
etcdctl --endpoints https://127.0.0.1:3379 \
  --cert /tmp/etcd-certs/certs/127.0.0.1.pem \
  --key /tmp/etcd-certs/certs/127.0.0.1-key.pem \
  --cacert /tmp/etcd-certs/certs/ca.pem \
  get a -wjson
```

### Log

```shell
tail -f /var/log/infra1.log
```


































