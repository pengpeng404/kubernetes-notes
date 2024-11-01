# code


## install env

```shell
git init
go mod init pp.com/pkg/pp-deployment
```

```shell
git add .
git commit -a
```


```shell
kubebuilder init --domain pp.com
kubebuilder create api --group apps --version v1 --kind PpDeployment
```


```shell
git add .
git commit -a
```

```shell
go mod tidy
go mod vendor
git add vendor
git commit -a
```

## 设计

```go
// PpDeploymentSpec defines the desired state of PpDeployment.
type PpDeploymentSpec struct {
// Image 镜像地址
Image string `json:"image"`
// Port 服务提供的端口
Port int32 `json:"port"`
// Replicas 要部署多少个副本
// +optional
Replicas int32 `json:"replicas,omitempty"`
// StartCmd 启动命令
// +optional
StartCmd string `json:"startCmd,omitempty"`
// Args 启动命令参数
// +optional
Args []string `json:"args,omitempty"`
// Environments 环境变量 直接使用 pod 中的定义方式
// +optional
Environments []corev1.EnvVar `json:"environments,omitempty"`
// Expose service 要暴露的端口
Expose *Expose `json:"expose"`
}

// Expose defines the desired state of Expose.
type Expose struct {
// Mode 模式 nodePort or ingress
Mode string `json:"mode"`
// IngressDomain 域名 在 Mode 为 ingress 时 必填
// +optional
IngressDomain string `json:"ingressDomain,omitempty"`
// NodePort nodePort 端口 在 Mode 为 nodePort 时 必填
// +optional
NodePort int32 `json:"nodePort,omitempty"`
// ServicePort service 端口 一般随机生成 这里使用和服务相同的端口 Port
// +optional
ServicePort int32 `json:"servicePort,omitempty"`
}



```


```go
// PpDeploymentStatus defines the observed state of PpDeployment.
type PpDeploymentStatus struct {
	// Phase 处于什么阶段
	// +optional
	Phase string `json:"phase,omitempty"`
	// Message 这个阶段信息
	// +optional
	Message string `json:"message,omitempty"`
	// Reason 处于这个阶段的原因
	// +optional
	Reason string `json:"reason,omitempty"`
	// Conditions 处于这个阶段的原因
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`
}

// Condition defines the observed state of Condition.
type Condition struct {
	// Type 子资源类型
	// +optional
	Type string `json:"type,omitempty"`
	// Message 子资源状态信息
	// +optional
	Message string `json:"message,omitempty"`
	// Status 子资源状态名称
	// +optional
	Status string `json:"status,omitempty"`
	// Reason 子资源状态原因
	// +optional
	Reason string `json:"reason,omitempty"`
	// LastTransitionTime 最后 创建/更新 时间
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}
```





























