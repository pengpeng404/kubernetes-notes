# CoreDNS


因为 pod 有生命周期，
podIP 会变化，但是 service 本身是个对象，
是固定的属性，创建出来就分好了，
相对稳定的访问入口，
ClusterIP 也是相对稳定，删除重建的话 IP 又变化，
所以不是一个真正百分百可靠，所以也不会把 serviceIP 给别人。

更常用的是把服务名称给出去，真正的服务发现 CoreDNS

CoreDNS
- 不落盘 in mem 效率高
- 本身有控制器 watch service endpoint pods 配置 DNS 记录
  - svc1.ns1.svc.clusterdomain:VIP1


客户端如何使用内置 DNS 服务

```shell
k run -it app --image centos -- /bin/bash
curl nginx-service.default.svc.cluster.local
#<!DOCTYPE html>
#<html>
#<head>
#<title>Welcome to nginx!</title>
#<style>
#html { color-scheme: light dark; }
#body { width: 35em; margin: 0 auto;
#font-family: Tahoma, Verdana, Arial, sans-serif; }
#</style>
#</head>
#<body>
#<h1>Welcome to nginx!</h1>
#<p>If you see this page, the nginx web server is successfully installed and
#working. Further configuration is required.</p>
#
#<p>For online documentation and support please refer to
#<a href="http://nginx.org/">nginx.org</a>.<br/>
#Commercial support is available at
#<a href="http://nginx.com/">nginx.com</a>.</p>
#
#<p><em>Thank you for using nginx.</em></p>
#</body>
#</html>

ping nginx-service.default.svc.cluster.local
#PING nginx-service.default.svc.cluster.local (10.100.129.52) 56(84) bytes of data.

#nginx-service   NodePort    10.100.129.52   <none>        80:30007/TCP   17h

cat /etc/resolv.conf 
#search default.svc.cluster.local svc.cluster.local cluster.local
#nameserver 10.96.0.10
#options ndots:5

ks get svc
#NAME       TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
#kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   22h

cat /etc/resolv.conf 
#search default.svc.cluster.local svc.cluster.local cluster.local
#nameserver 10.96.0.10
#options ndots:5
# 查询 FQDN 的点数 然后从 search 中一个一个 append 然后再 DNS 查询

```


```shell
k get po app -oyaml
#dnsPolicy: ClusterFirst
# pod 内部 /etc/resolv.conf
```

```yaml
# 也可以自定义 DNS server 做域名查询
spec:
  dnsPolicy: "None"
  dnsConfig:
    nameservers:
      - 1.2.3.4
    searchs:
      - xx.ns1.svc.cluster.local
      - xx.daemon.com
    options:
      - name: ndots
        values: "2"
```



### 不同类型服务 DNS 记录

- 普通 Service
  - ClusterIP nodePort LoadBalancer 都有 apiserver 分配的 ClusterIP
  - FQDN svc1.ns1.svc.clusterdomain:ClusterIP
- Headless Service
  - ClusterIP none 没有虚 IP
  - 被 CoreDNS 监听
  - FQDN podname1.ns1.svc.clusterdomain:PodIP(ready)
  - statefulSet 每个 pod 都有一个独立的域名
  - 本身 statefulSet 名字是递增的 不一样
- ExternalName Service
  - 引用一个已经存在的域名
  - CoreDNS 为该 Service 创建一个 CName 记录 指向目标域名



### DNS 落地实践

DNS 查询服务本身不是做负载均衡的，




















































