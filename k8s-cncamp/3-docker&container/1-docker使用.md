
```shell
docker run -it centos bash
```

如果本地没有这个镜像文件，会运行
```shell
docker pull centos
```
从 dockerhub 镜像仓库拉下来 
，实际上就是一个tar包加描述文件
，然后解压这个tar包，把这个解压的文件作为进程的 rootfs
，让这个进程起来
，起来的这个进程就是 bash

```shell
ps -ef|grep bash
#root        1670    1309  0 11:11 pts/0    00:00:00 docker run -it centos:7.4.1708 bash
```

## docker 命令
```shell
# 查询当前运行了哪些容器
docker ps
# 查看容器细节
docker inspect <containerid>
```



