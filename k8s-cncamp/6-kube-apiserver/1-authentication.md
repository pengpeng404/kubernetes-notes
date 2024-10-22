# 认证

开启 TLS 时，所有的请求都需要首先认证。

Kubernetes 支持多种认证机制，并支持同时开启多个认证插件，只要有一个认证通过即可。
- 认证成功 -> 授权验证
- 认证失败 -> 返回 `http401`


## 认证插件

```shell
cat ~/.kube/config
# 认证信息
```

### X509证书

```shell
# Create private key and csr
openssl genrsa -out ppnew.key 2048
openssl req -new -key ppnew.key -out ppnew.csr -subj "/CN=pppp/O=group1"

# 提交 CSR 到 Kubernetes
kubectl apply -f - <<EOF
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: ppnew
spec:
  request: $(cat ppnew.csr | base64 | tr -d '\n')
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - client auth
EOF

kubectl certificate approve ppnew
kubectl get csr ppnew -o jsonpath='{.status.certificate}' | base64 --decode > ppnew.crt

# 配置用户凭据
kubectl config set-credentials pppp --client-certificate=ppnew.crt --client-key=ppnew.key

k get po --user pppp
# 此时还没有给定权限 但是用户已经认证成功是 pppp
#Error from server (Forbidden): pods is forbidden: User "pppp" cannot list resource "pods" in API group "" in the namespace "default"

# Grant permission
kubectl create role pp-role --verb=create --verb=get --verb=list --verb=update --verb=delete --resource=pods --namespace=default
kubectl create rolebinding developer-binding-myuser --role=pp-role --user=pppp

k get po --user pppp
#NAME                       READY   STATUS             RESTARTS      AGE
#apphttp-5887bfd99d-ncgwx   1/1     Running            2 (23h ago)   8d
#my-release-etcd-0          0/1     ImagePullBackOff   0             17h

ks get po --user pppp
#Error from server (Forbidden): pods is forbidden: User "pppp" cannot list resource "pods" in API group "" in the namespace "kube-system"
# 没有给 kube-system 的权限
```




### 静态Token文件
```shell
# Put static-token to target folder
mkdir -p /etc/kubernetes/auth
cp 1-1-static-token.csv /etc/kubernetes/auth/

# Backup your orginal apiserver /root/
cp /etc/kubernetes/manifests/kube-apiserver.yaml ~/kube-apiserver.yaml

# Override your kube-apiserver with the one with static-token config
cp kube-apiserver.yaml /etc/kubernetes/manifests/kube-apiserver.yaml

# Get kubernetes object with static token
curl https://192.168.34.2:6443/api/v1/namespaces/default -H "Authorization: Bearer pp-token" -k

#{
#  "kind": "Status",
#  "apiVersion": "v1",
#  "metadata": {},
#  "status": "Failure",
#  "message": "namespaces \"default\" is forbidden: User \"pp\" cannot get resource \"namespaces\" in API group \"\" in the namespace \"default\"",
#  "reason": "Forbidden",
#  "details": {
#    "name": "default",
#    "kind": "namespaces"
#  },
#  "code": 403
#}

vi ~/.kbue/config
#- name: pp
#  user:
#    token: pp-token
k get po --user pp
#Error from server (Forbidden): pods is forbidden: User "pp" cannot list resource "pods" in API group "" in the namespace "default"
# 因为此时还没有授权
```


### ServiceAccount

每个 `Namespace` 都有一个默认的 `default` `ServiceAccount`

`ServiceAccount` 为 `Pod` 提供一个访问 `Kubernetes API` 的身份

当我的 `pod` 不与 `apiServer` 交互 则无关紧要

```shell
k get po apphttp-5887bfd99d-ncgwx -oyaml

#volumeMounts:
#- mountPath: /var/run/secrets/kubernetes.io/serviceaccount
#  name: kube-api-access-49z8d
#  readOnly: true
#  
#serviceAccount: default
#serviceAccountName: default
```

`ServiceAccount` 应用场景
- `operator` 需要与 `apiServer` 交互 修改数据是否有权限
- 用户认证 直接生成 `ServiceAccount` 直接使用 `ServiceAccount token`

### webhoook 令牌身份认证

每个公司都有自己的认证信息，
不需要为自己的 `kubernetes` 重新开发认证平台。
通过 `webhook` 把已经有的认证平台加入到 `kubernetes`

`webhook` 只有 `tokenReview` 方式

```log
/authn-webhook
在这个目录文件夹下展示了如何使用 github token 的外部认证平台
来进行 kubernetes 认证
```

自己的认证系统大概率不能适配 `kubernetes` 的 `webhook` 标准

所以需要自己实现适配器 参考 `/authn-webhook`




## 认证系统在生产系统中遇到的陷阱

基于 `KeyStone` 的认证插件导致 `KeyStone` 故障且无法恢复

`KeyStone` 是企业关键服务 并且 `Kubernetes` 以 `KeyStone` 作为认证插件

`KeyStone` 出现故障后会抛出 `401` 错误

`Kubernetes` 发现 `401` 错误后会尝试重新认证

（任何的控制器都应该遵循指数级重试）

大多数的控制器都有指数级 `back off` 重试间隔越来越慢

但是 `gopherCloud` 针对过期的 `token` 会一直 `retry`

大量的 `request` 积压在 `KeyStone` 导致服务无法恢复

`Kubernetes` 成为压死企业认证服务的最后一根稻草

解决方案：
- Circuit break 熔断
- Rate limit






















































