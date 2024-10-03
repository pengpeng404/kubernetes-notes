# 构建 httpserver 容器镜像

### httpserver

[httpserver](../0-httpserver/0-easy-httpserver/main.go)


将 Go 程序编译成二进制文件 在windows操作系统上略有不同

```shell
GOOS=linux GOARCH=amd64 go build -o httpserver
```


### 构建镜像

Dockerfile

```dockerfile
FROM alpine:3.2
WORKDIR /app
ADD httpserver /app/httpserver
EXPOSE 80
ENTRYPOINT ["/app/httpserver"]
```

```shell
# -t myhttpserver 表示镜像名称
# v1.0 版本号
# . 从当前目录构建
docker build -t apphttp:v1.0 .
```

### 启动容器镜像

```shell
docker run -d -p 8888:80 apphttp:v1.0
```

### 排查

```shell
docker logs xxxxxxxx
```

### 访问服务
```
192.168.34.110:8888/?user=pp
192.168.34.110:8888/healthz
```













