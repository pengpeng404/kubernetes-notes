# 手动配置docker网络


### Create network ns

```sh
mkdir -p /var/run/netns
# 删除无效的符号链接
# 在 /var/run/netns/ 目录中，
# 符号链接通常指向进程的网络命名空间（如 /proc/$pid/ns/net）。
# 当进程结束时，对应的 /proc/$pid/ns/net 目录会被内核清理掉，
# 而符号链接则可能会变成“悬空”状态，指向无效的路径。
# 执行 find -L 的目的是找到那些指向无效目标的符号链接，
# 并将它们删除，防止后续操作中使用到无效的命名空间。
find -L /var/run/netns -type l -delete
```

### Start nginx docker with non network mode

```sh
# -d 后台运行
docker run --network=none  -d nginx:1.27.0
```

### Check corresponding pid

```sh
docker ps|grep nginx
docker inspect <containerid>|grep -i pid

"Pid": 1052036,
"PidMode": "",
"PidsLimit": null,
```

### Check network config for the container

```sh
nsenter -t 1052036 -n ip a
```
```log
root@docker:/home/pp# nsenter -t 1052036 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
```


### Link network namespace

```sh
export pid=1052036
# 创建符号链接，
# 将某个进程的网络命名空间关联到 /var/run/netns/ 目录中，
# 以便可以通过 ip netns 工具来管理该进程的网络命名空间
ln -s /proc/$pid/ns/net /var/run/netns/$pid
ip netns list
```

### Check docker bridge on the host

```sh
brctl show

root@docker:/home/pp# brctl show
bridge name     bridge id               STP enabled     interfaces
docker0         8000.0242264abaf5       no
```
```shell
ip a

4: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:26:4a:ba:f5 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:26ff:fe4a:baf5/64 scope link 
       valid_lft forever preferred_lft forever
```

### Create veth pair

```sh
ip link add A type veth peer name B
```

### Config A

```sh
# 将虚拟网络设备 A 接入桥接设备 docker0
brctl addif docker0 A
# 开启
ip link set A up
```

### Config B

```sh
SETIP=172.17.0.15
SETMASK=16
GATEWAY=172.17.0.1

# 将网络接口 B 移动到指定的网络命名空间
ip link set B netns $pid
```
```shell
# 将网络接口 B 重命名为 eth0
ip netns exec $pid ip link set dev B name eth0
```
```log
此时可以看到该容器的网络配置已经改变

root@docker:/home/pp# nsenter -t 1052036 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
13: eth0@if14: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 06:a9:5f:ac:55:cf brd ff:ff:ff:ff:ff:ff link-netnsid 0
```

```shell
# 开启
ip netns exec $pid ip link set eth0 up
```
```log
root@docker:/home/pp# nsenter -t 1052036 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
13: eth0@if14: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 06:a9:5f:ac:55:cf brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet6 fe80::4a9:5fff:feac:55cf/64 scope link 
       valid_lft forever preferred_lft forever
```

```shell
# 在指定的网络命名空间中，为网络接口 eth0 配置 IP 地址
ip netns exec $pid ip addr add $SETIP/$SETMASK dev eth0
```

```log
root@docker:/home/pp# nsenter -t 1052036 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
13: eth0@if14: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 06:a9:5f:ac:55:cf brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.15/16 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::4a9:5fff:feac:55cf/64 scope link 
       valid_lft forever preferred_lft forever
```


```shell
# 在指定的网络命名空间中添加一条默认路由
ip netns exec $pid ip route add default via $GATEWAY
```


```log
root@docker:/home/pp# curl 172.17.0.15
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

如果需要做端口转发，手动配置 `iptables` 重启容器以上配置丢失
