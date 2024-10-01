# Namespace

系统可以为进程分配不同的 Namespace

保证不同的 Namespace 资源独立分配 进程彼此隔离

## Linux Namespace 常用操作
```shell
# 查看当前系统的 ns
lsns -t <type>
# 查看某进程的 ns 可以查看指定进程的各种命名空间（namespace）实例
ls -la /proc/<pid>/ns/
# 进入某个 ns 执行命令
nsenter -t <pid> -n ip addr
```

## 查看运行时的 Namespace

```shell
# step 1 找到运行的容器
docker ps|grep <container name>
# step 2 查看 spec 找到主机 pid
docker inspect <CONTAINER ID>|grep -i pid
# step 3 使用 nsenter 运行命令 这里查看网络配置
nsenter -t <pid> -n ip a
```

注意：使用 `ps -ef` 查看的 `docker run` 是守护进程，用来管理
容器的生命周期，而不是容器内的实际应用进程。这个守护进程是在主机的
命名空间中。

而容器内部进程的真实 `pid` 需要通过 `docker inspect` 查看。

## Namespace 练习

```shell
# step 1 unshare -fn 命令创建了一个新的网络命名空间，并让 unshare 自己进入这个新的命名空间
# 然后 unshare 会 fork 出一个新的子进程 sleep，这个 sleep 进程继承了父进程（unshare）的命名空间，
# 因此 sleep 和 unshare 进程都在同一个新的网络命名空间中，而不再是宿主机的网络命名空间。
# sleep 是 unshare 的子进程，它继承了父进程的命名空间。
# 因此，当你通过 nsenter 进入 unshare 或 sleep 的 PID 时，
# 实际上你进入的是同一个新创建的网络命名空间。
unshare -fn sleep 180
# step 2 找到 unshare 进程和 sleep 进程的 pid
ps -ef|grep sleep
# step 3 查看两个进程的网络配置
nsenter -t <pid> -n ip a
# step 4 查看 ns
ls -la /proc/<pid>/ns/
```

















