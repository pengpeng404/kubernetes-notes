



```shell
mkdir -p /tmp/etcd-download-test

tar xzvf etcd-xxx.tar.gz -C /tmp/etcd-download-test --strip-components=1

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
/tmp/etcd-download-test/etcdutl version

sudo mv /tmp/etcd-download-test/etcd /usr/local/bin/
sudo mv /tmp/etcd-download-test/etcdctl /usr/local/bin/
sudo mv /tmp/etcd-download-test/etcdutl /usr/local/bin/
sudo chmod +x /usr/local/bin/etcd /usr/local/bin/etcdctl /usr/local/bin/etcdutl


```


























