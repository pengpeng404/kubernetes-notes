
# 代码组织结构

## 生成 CRD 代码

```shell
mkdir -p /home/pp/repos/demo
cd /home/pp/repos/demo

go mod init pengpeng.com/demo/app
```

```shell
# 初始化 kubebuilder
## 这将生成一个基础的项目结构，并包含必要的文件和目录
kubebuilder init --domain pengpeng.com
## group 组 同一类资源集合
## version 版本信息
## kind 首字母大写 驼峰
kubebuilder create api --group demo --version v1 --kind App
## 这将生成以下内容
### api/demo/v1/app_types.go：定义了 App 资源的结构
### controllers/app_controller.go：控制器代码框架
```

```shell
make generate
make manifests
# make generate 会生成 zz_generated.deepcopy.go 文件，用于对象的深拷贝方法
# make manifests 会在 config/crd 目录下生成 CRD YAML 文件
```



### 流程

#### step 1 生成 CRD controller 代码
```shell
kubebuilder init --domain pengpeng.com
kubebuilder create api --group demo --version v1 --kind App
```

#### step 2 定义资源

```shell
# Spec
# Status
```

#### step 3 生成代码和清单文件

```shell
make generate
make manifests
# 用来生成操作对象的工具 以及生成对象的 yaml 文件
# CRD 清单文件和 RBAC 权限配置
```

#### step 4 部署 CRD 到集群

```shell
make install
# 查看是否成功创建
kubectl get crds
```

#### step 5 编译运行

```shell
make build
#go fmt ./...
#go vet ./...
#go build -o bin/manager cmd/main.go

make run
#go run ./cmd/main.go
# 运行 controller 二进制文件
```

#### step 6 创建实例并验证 controller


## 实操

```shell
# 一旦修改了 api 文件
# 都要重新执行一下命令
make generate
make manifests
# 如果字段中有其他类型 还需要进行相关的修改 copy 函数

#pp@DESKTOP-HNF938B:~/repos/demo$ make generate
#/home/pp/repos/demo/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
#pp@DESKTOP-HNF938B:~/repos/demo$ make manifests
#/home/pp/repos/demo/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

```

执行之后查看 `config/crd/bases/xxxx.yaml` 是否符合预期

```shell
# 把 CRD 加入到 k8s
make install
#/home/pp/repos/demo/bin/kustomize build config/crd | kubectl apply -f -
#customresourcedefinition.apiextensions.k8s.io/apps.demo.pengpeng.com created

```


```shell
# 只是把对象信息建立出来 目前还没有实例
k get crd apps.demo.pengpeng.com -oyaml
```

```shell
# 执行 controller 二进制文件
make build
#go fmt ./...
#go vet ./...
#go build -o bin/manager cmd/main.go

make run
#go run ./cmd/main.go
# 运行 controller 二进制文件
```


```yaml
# demo.yaml
apiVersion: demo.pengpeng.com/v1
kind: App
metadata:
  labels:
    app.kubernetes.io/name: demo
    app.kubernetes.io/managed-by: kustomize
  name: app-sample
spec:
  action: hello
  object: kubernetes

```

```shell
k create -f demo.yaml
#app.demo.pengpeng.com/app-sample created

k get apps.demo.pengpeng.com
#NAME         AGE
#app-sample   8m19s

k get apps.demo.pengpeng.com app-sample -oyaml
#  spec:
#    action: hello
#    object: kubernetes
#  status:
#    result: hello++kubernetes

# 和控制器写的一样
```

```shell
# make run
#start app reconciliation
#handle data
#End app reconciliation
#start app reconciliation
#handle data
#End app reconciliation

# 一共触发两次 Reconcile
# 做了一次更新 又触发一次 Reconcile
```


## Operator 开发代码


更新对象之前进行深度拷贝 更新拷贝对象 这样对缓存没有污染


```go
// AppSpec defines the desired state of App.
type AppSpec struct {
// Action
//+optional
Action string `json:"action,omitempty"`
// Object
//+optional
Object string `json:"object,omitempty"`
}

// AppStatus defines the observed state of App.
type AppStatus struct {
// Result
//+optional
Result string `json:"result,omitempty"`
}

```


```go
// Reconcile
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("start app reconciliation")
	// 写自己的逻辑

	// 1 拿到这个对象
	app := new(demov1.App)
	if err := r.Client.Get(ctx, req.NamespacedName, app); err != nil {
		// 如果找不到这个对象 则返回一个非nil 否则会一直循环
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// 2 处理数据
	logger.Info("handle data")
	action := app.Spec.Action
	object := app.Spec.Object

	result := fmt.Sprintf("%s++%s", action, object)

	// 3 创建结构
	appCopy := app.DeepCopy()
	appCopy.Status.Result = result

	// 4 更新结果
	err := r.Client.Status().Update(ctx, appCopy)
	if err != nil {
		return ctrl.Result{}, err
	}

	logger.Info("End app reconciliation")

	return ctrl.Result{}, nil
}

```





```shell
k get crds
k get <crd-name>
k delete <crd-name> <crd-resource-name>



```









































