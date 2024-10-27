# VirtualBox Kubernetes 1.30.4 containerd calico

三台虚拟机 1 master + 2 node

- 192.168.34.101 k8s-master
- 192.168.34.102 k8s-node1
- 192.168.34.103 k8s-node2

## 虚拟机安装 linux(ubuntu 20.04 server)

https://www.virtualbox.org/wiki/Downloads

https://releases.ubuntu.com/20.04/

- NAT 网络 VirtualBox 自动分配
- Host-Only 192.168.34.1 255.255.255.0 禁用 DHCP 服务器
- 至少2核4G

---

- step 1 安装 linux 系统 安装过程中 选择安装 SSH 工具 其余默认选项 安装完成后 关机
- step 2 配置 VirtualBox 的网络适配器（设置网卡 1 为 NAT 网络 2 为 Host-Only）
- step 3 进入虚拟机并编辑 /etc/netplan/00-installer-config.yaml 来配置虚拟机的网络接口 IP
- step 4 保存并应用网络配置，重启网络服务后即可使用

```shell
# sudo su 进入 root 模式
# nano ctrl + o enter 保存修改 ctrl + x 退出
nano /etc/netplan/00-installer-config.yaml
# 每个主机单独配置 IP
network:
  ethernets:
    enp0s3:
      dhcp4: true
    enp0s8:
      dhcp4: no
      addresses:
        - 192.168.34.101/24
  version: 2

netplan apply
```

```shell
# 更新包列表并安装 chrony
sudo apt update && sudo apt install -y chrony
```

```shell
# 备份原始的配置文件
sudo cp /etc/chrony/chrony.conf /etc/chrony/chrony.conf.bak
# 写入 server ntp.aliyun.com 配置到 chrony.conf
bash -c 'echo "server ntp.aliyun.com iburst" >> /etc/chrony/chrony.conf'
# 启动并启用 chrony 服务
sudo systemctl start chrony
sudo systemctl enable chrony
```

```shell
# 打印 chrony 状态
sudo systemctl status chrony
```

```shell
# 验证同步状态
timedatectl
date
```

## Install kubernetes

### 关闭防火墙

```sh
service ufw stop
update-rc.d ufw defaults-disabled

swapoff -a
sed -ri 's/.*swap.*/#&/' /etc/fstab
```


### 系统优化

```shell
sudo cat > /etc/sysctl.d/k8s_better.conf << EOF
net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-ip6tables=1
net.ipv4.ip_forward=1
vm.swappiness=0
vm.overcommit_memory=1
vm.panic_on_oom=0
fs.inotify.max_user_instances=8192
fs.inotify.max_user_watches=1048576
fs.file-max=52706963
fs.nr_open=52706963
net.ipv6.conf.all.disable_ipv6=1
net.netfilter.nf_conntrack_max=2310720
EOF

modprobe br_netfilter
lsmod | grep conntrack
modprobe nf_conntrack
sysctl -p /etc/sysctl.d/k8s_better.conf
```

### 开启 IPVS
```sh
# 安装依赖包
apt-get install -y conntrack ipvsadm ipset jq iptables curl sysstat wget vim net-tools git
```

```shell
## 开启 IPVS 转发
modprobe br_netfilter

mkdir -p /etc/sysconfig/modules/
cat > /etc/sysconfig/modules/ipvs.modules << EOF
#!/bin/bash
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack
EOF

chmod 755 /etc/sysconfig/modules/ipvs.modules

bash /etc/sysconfig/modules/ipvs.modules

lsmod | grep -e ip_vs -e nf_conntarck
```

### containerd

https://github.com/containerd/containerd/releases/download/v1.7.20/cri-containerd-cni-1.7.20-linux-amd64.tar.gz

```sh
# 解压根目录
tar -zxvf xxxxx -C /
```

```shell
# 创建配置目录
mkdir -p /etc/containerd
# 生成默认配置文件
containerd config default > /etc/containerd/config.toml
```

```shell
# 使用 nano 编辑器进行手动修改
sudo nano /etc/containerd/config.toml

# 将以下内容：
SystemdCgroup = false 
# 修改为：
SystemdCgroup = true

# 将以下内容：
sandbox_image = "k8s.gcr.io/pause:3.6"
# 修改为：
sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.9"

# 修改容器镜像仓库
[plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
  endpoint = ["hwcloud"]
  
/etc/containerd/certs.d
cd /etc/containerd/
mkdir -p certs.d && cd certs.d/
sudo systemctl restart containerd

# docker hub镜像加速
mkdir -p /etc/containerd/certs.d/docker.io
cat > /etc/containerd/certs.d/docker.io/hosts.toml << EOF
server = "https://docker.io"
[host."https://dockerproxy.cn"]
  capabilities = ["pull", "resolve"]

[host."https://docker.m.daocloud.io"]
  capabilities = ["pull", "resolve"]
EOF

# registry.k8s.io镜像加速
mkdir -p /etc/containerd/certs.d/registry.k8s.io
tee /etc/containerd/certs.d/registry.k8s.io/hosts.toml << 'EOF'
server = "https://registry.k8s.io"

[host."https://k8s.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# docker.elastic.co镜像加速
mkdir -p /etc/containerd/certs.d/docker.elastic.co
tee /etc/containerd/certs.d/docker.elastic.co/hosts.toml << 'EOF'
server = "https://docker.elastic.co"

[host."https://elastic.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# gcr.io镜像加速
mkdir -p /etc/containerd/certs.d/gcr.io
tee /etc/containerd/certs.d/gcr.io/hosts.toml << 'EOF'
server = "https://gcr.io"

[host."https://gcr.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# ghcr.io镜像加速
mkdir -p /etc/containerd/certs.d/ghcr.io
tee /etc/containerd/certs.d/ghcr.io/hosts.toml << 'EOF'
server = "https://ghcr.io"

[host."https://ghcr.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# k8s.gcr.io镜像加速
mkdir -p /etc/containerd/certs.d/k8s.gcr.io
tee /etc/containerd/certs.d/k8s.gcr.io/hosts.toml << 'EOF'
server = "https://k8s.gcr.io"

[host."https://k8s-gcr.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# mcr.m.daocloud.io镜像加速
mkdir -p /etc/containerd/certs.d/mcr.microsoft.com
tee /etc/containerd/certs.d/mcr.microsoft.com/hosts.toml << 'EOF'
server = "https://mcr.microsoft.com"

[host."https://mcr.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# nvcr.io镜像加速
mkdir -p /etc/containerd/certs.d/nvcr.io
tee /etc/containerd/certs.d/nvcr.io/hosts.toml << 'EOF'
server = "https://nvcr.io"

[host."https://nvcr.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# quay.io镜像加速
mkdir -p /etc/containerd/certs.d/quay.io
tee /etc/containerd/certs.d/quay.io/hosts.toml << 'EOF'
server = "https://quay.io"

[host."https://quay.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# registry.jujucharms.com镜像加速
mkdir -p /etc/containerd/certs.d/registry.jujucharms.com
tee /etc/containerd/certs.d/registry.jujucharms.com/hosts.toml << 'EOF'
server = "https://registry.jujucharms.com"

[host."https://jujucharms.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF

# rocks.canonical.com镜像加速
mkdir -p /etc/containerd/certs.d/rocks.canonical.com
tee /etc/containerd/certs.d/rocks.canonical.com/hosts.toml << 'EOF'
server = "https://rocks.canonical.com"

[host."https://rocks-canonical.m.daocloud.io"]
  capabilities = ["pull", "resolve", "push"]
EOF
```

```shell
# 设置 containerd 开机启动并立即启动服务
systemctl enable containerd
systemctl start containerd

# 重启
sudo systemctl restart containerd


# 列出 containerd 管理的镜像
ctr images ls
# 查看 runc 版本
runc --version
```

### 配置 overlay 转发
```sh
# 创建 /etc/modules-load.d/containerd.conf 文件并写入内容
cat << EOF > /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF

# 手动加载 overlay 和 br_netfilter 模块
sudo modprobe overlay
sudo modprobe br_netfilter

# 验证模块是否正确加载
lsmod | grep overlay
lsmod | grep br_netfilter
```


### 安装 k8s 1.30 aliyun


```sh
# 根据阿里云镜像仓库的步骤来
sudo mkdir -p /etc/apt/keyrings
apt-get update && apt-get install -y apt-transport-https
```

```shell
curl -fsSL https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.30/deb/Release.key | \
    gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.30/deb/ /" | \
    tee /etc/apt/sources.list.d/kubernetes.list
```

```shell
apt-get update
```

```shell
apt-get install -y kubelet kubeadm kubectl
```

```shell
# 验证 K8S 版本
kubelet --version
kubectl version --client

# 关闭自动更新
apt-mark hold kubelet kubeadm kubectl
```

### Cgroup

```sh
# 配置 kubelet 使用 systemd 作为 cgroup 驱动
sudo bash -c 'echo KUBELET_EXTRA_ARGS="--cgroup-driver=systemd" > /etc/sysconfig/kubelet'

# 设置 kubelet 开机自启
sudo systemctl enable kubelet

# 启动 kubelet 服务
sudo systemctl start kubelet
```

```sh
# 查看 K8S 所需要的镜像
kubeadm config images list --kubernetes-version=v1.30.6
# 拉取镜像
kubeadm config images pull --image-repository registry.aliyuncs.com/google_containers
```

```shell
# K8S 初始化 只有 master 节点初始化 其他节点加入 master
kubeadm init --kubernetes-version=v1.30.4 --pod-network-cidr=10.224.0.0/16 --apiserver-advertise-address=192.168.34.101 --image-repository registry.aliyuncs.com/google_containers --ignore-preflight-errors=all

mkdir -p $HOME/.kube
sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

export KUBECONFIG=/etc/kubernetes/admin.conf

# node 节点加入 master
```

### 安装 calico

https://raw.githubusercontent.com/projectcalico/calico/v3.27.3/manifests/calico.yaml

```sh
# 注意 yaml 文件格式
# 修改 calico
    # - name: CALICO_IPV4POOL_CIDR
    #   value: "192.168.0.0/16"
    - name: CALICO_IPV4POOL_CIDR
      value: "10.244.0.0/16"
    - name: IP_AUTODETECTION_METHOD
      value: "interface=enp0s8"

# 查看所需镜像
cat calico.yaml | grep image
## 对应版本下载镜像
## https://docker.aityp.com/

kubectl apply -f calico.yaml
kubectl get pod -n kube-system -o wide
kubectl get node -A
```


```shell
ctr -n k8s.io images pull
ctr -n k8s.io images tag

ctr -n k8s.io images import my-image.tar

#ctr -n k8s.io images pull swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/nicolaka/netshoot:v0.13
#ctr -n k8s.io images tag swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/nicolaka/netshoot:v0.13 docker.io/nicolaka/netshoot:v0.13
#ctr -n k8s.io images tag docker.io/nicolaka/netshoot:v0.13 docker.io/library/nicolaka/netshoot:v0.1

```


```shell
visudo
#cadmin ALL=(ALL) NOPASSWD:ALL

sudo nano ~/.bashrc
#alias k="kubectl"
#alias ks="kubectl -n kube-system"
source ~/.bashrc
```























