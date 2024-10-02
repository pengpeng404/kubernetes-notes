# Dockerfile最佳实践

## 最佳实践
- 目标：易管理、少漏洞、镜像小、层级少、利用缓存
- 1、不要安装无效软件包
- 2、每个镜像应该只有一个进程
- 3、无法避免同一镜像运行多进程时，应选择合理的初始化进程
    - 需要捕获 `SIGTERM` 信号并完成子进程的优雅终止
    - 负责清理推出的子进程 避免僵尸进程
- 4、最小化层级
  - 多条 `RUN` 指令通过连接符连接成一条指令 减少层数
  - 多段构建减少镜像层数
- 5、多行参数按照字母排序
- 6、编写 `Dockerfile` 有效利用 `build cache`
- 7、复制文件时 每个文件单独复制 确保某个文件变更时 只影响该文件对应的缓存

如果一个容器需要多个进程并行的时候，
可以使用 `tini` 作为初始化进程

https://github.com/krallin/tini

## 构建上下文 Build Context

运行 `docker build` 的时候，当前工作目录被称为构建构建上下文，
需要把当前目录中的内容发送到 `docker daemon`，
当后面执行 `Dockerfile` 中的每条指令的时候，
就从传输的上下文里寻找，
比如通过 `ADD` 要加一个文件，
这个文件必须在 `Dockerfile` 所在目录存在。

如果 `Dockerfile` 所在的根目录文件很大，
那么构建上下文就会很大，
构建镜像会很久很久，
因为需要把构建上下文全都传给`docker daemon`。

可以通过 `.dockerignore` 文件从编译上下文中排除某些文件。


## Build Cache

最佳实践：变动不频繁的层级放在 `Dockerfile` 上面，
变动频繁的放在 `Dockerfile` 下面。


## 多段构建

有效减少镜像层级的方式

```dockerfile
# 构建一个带有 Go 编译器的基础镜像
# 安装必要依赖并编译项目
# 生成二进制文件
FROM golang:1.16-alpine AS build
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
RUN dep ensure --vendor-only

COPY . /go/src/project/
RUN go build -o /bin/project

# 使用 FROM scratch 构建最终的镜像
# 仅包含编译好的二进制文件
# 避免将不必要的编译工具和依赖带入最终的生产环境
FROM scratch
COPY --from=build /bin/project /bin/project
ENTRYPOINT ["/bin/project"]
CMD ["--help"]
```


## Dockerfile常用指令

- FROM 选择基础镜像
```dockerfile
FROM [--platform=<platform>] <image>[@<digest>] [AS <name>]
```

- LABEL 标签 可以配合 `label filter` 命令过滤镜像查询结果
```dockerfile
LABEL multi.label1="value1" multi.label2="value2"
```
```shell
docker images -f label=multi.label1="value1"
```

- RUN 执行命令
```dockerfile
RUN apt-get update && apt-get install -y <package>
```
注意：这两条命令应该永远使用 `&&` 链接，
如果分开执行，
`RUN apt-get update` 构建层会被缓存，
可能导致新的包无法安装。

- CMD 定义容器镜像中包含的运行命令，通常用于应用启动
```dockerfile
CMD ["nginx", "-g", "daemon off;"]
```
当 `CMD` 和 `ENTRYPOINT` 一起使用时（最佳实践），
`CMD` 的作用是为 `ENTRYPOINT` 提供默认参数
```dockerfile
ENTRYPOINT ["python"]
CMD ["app.py"]
```

- ENV 设置环境变量
```dockerfile
ENV <key>=<value>
```

- ADD 从源地址（文件，目录或者 URL）复制文件到目标路径
```dockerfile
ADD [--chown=<user>:<group>] <src>... <dest>`
ADD [--chown=<user>:<group>] ["<src>", ... "<dest>"]`
```
`src` 如果是文件，
则必须包含在编译上下文中，
`ADD` 指令无法添加编译上下文之外的文件

`src` 如果是一个目录，
则所有文件都会被复制至 `dest`

`src` 如果是一个本地压缩文件，
则在 `ADD` 的同时完整解压操作

如果 `dest` 不存在，
则 `ADD` 指令会创建目标目录

应尽量减少通过 `ADD URL` 添加 remote 文件，
建议使用 `curl` 或者 `wget && untar`

- COPY 可以用于多段编译 在前一个临时镜像中拷贝文件
```dockerfile
COPY --from=build /bin/project /bin/project
```

















