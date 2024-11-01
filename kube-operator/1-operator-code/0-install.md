
# 环境配置

```log
使用 WSL 开发 有好几个坑

```

### WSL

```shell
# windows
# 查看当前环境 wls
wsl --list
# 卸载
wsl --unregister Ubuntu
wsl --list
# 查看可安装的 Linux 子系统
wsl --list --online

wsl --install
```


### docker

```shell
# General -> Use the WSL 2 based engine
# Resources/WSL integration -> Enable Ubuntu

docker version
#Client:
# Version:           27.2.0
# API version:       1.47
# Go version:        go1.21.13
# Git commit:        3ab4256
# Built:             Tue Aug 27 14:14:20 2024
# OS/Arch:           linux/amd64
# Context:           default
#
#Server: Docker Desktop  ()
# Engine:
#  Version:          27.2.0
#  API version:      1.47 (minimum version 1.24)
#  Go version:       go1.21.13
#  Git commit:       3ab5c7d
#  Built:            Tue Aug 27 14:15:15 2024
#  OS/Arch:          linux/amd64
#  Experimental:     false
# containerd:
#  Version:          1.7.20
#  GitCommit:        8fc6bcff51318944179630522a095cc9dbf9f353
# runc:
#  Version:          1.1.13
#  GitCommit:        v1.1.13-0-g58aa920
# docker-init:
#  Version:          0.19.0
#  GitCommit:        de40ad0
```

```log
按照这样设置之后 打开 docker desktop
就可以在 WSL 中使用 docker
docker ps
```

```shell
# tools
sudo apt update
sudo apt install build-essential
```

### golang
```shell
# install go
sudo su
tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.bashrc
source ~/.bashrc
echo "export GOPATH=/root/go" >> ~/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc

go env GOPROXY
go env GOPATH

# install go NonRoot
mkdir -p go
tar -C go -xzf go1.20.5.linux-amd64.tar.gz --strip-components=1

nano ~/.bashrc
######################################################################
# 设置 GOROOT 为 Go 的安装路径
export GOROOT=/home/pp/go

# 添加 GOROOT/bin 到 PATH 中
export PATH=$PATH:$GOROOT/bin

# 设置 GOPATH 为 Go 工作空间（可选）
export GOPATH=/home/pp/pkg
export PATH=$PATH:$GOPATH/bin

export GOPROXY=https://goproxy.cn,direct
######################################################################
source ~/.bashrc


```

### kubectl kubebuilder kind

```shell
# install kubectl
## https://dl.k8s.io/release/v1.31.0/bin/linux/amd64/kubectl
sudo chmod +x kubectl
sudo mv kubectl /usr/local/bin

# install kind
## https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64
chmod +x kind
sudo mv kind /usr/local/bin

# 可以现在 docker 上 pull 下来 在 WSL 中可以看到
kind create cluster --image kindest/node:v1.31.1
#Creating cluster "kind" ...
# ✓ Ensuring node image (kindest/node:v1.31.1) 🖼
# ✓ Preparing nodes 📦  
# ✓ Writing configuration 📜 
# ✓ Starting control-plane 🕹️ 
# ✓ Installing CNI 🔌 
# ✓ Installing StorageClass 💾 
#Set kubectl context to "kind-kind"
#You can now use your cluster with:
#
#kubectl cluster-info --context kind-kind
kubectl cluster-info --context kind-kind

# 删除默认集群
kind delete cluster

```


```shell
# install kubebuilder
sudo chmod +x kubebuilder
sudo mv kubebuilder /usr/local/bin
```






































