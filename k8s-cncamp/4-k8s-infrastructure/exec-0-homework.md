



- 启动一个 EnvoyDeployment
- 要求 Envoy 的启动配置从外部的配置文件 Mount 到 Pod
- 进入 Pod 查看 Envoy 进程和配置
- 更改配置的监听端口并测试访问入口的变化
- 通过非级联的方式逐个删除对象




```shell
k create configmap envoy-config --from-file=envoy.yaml=ks-envoy-config.yaml
k create -f ks-envoy-deployment.yaml
k expose deploy envoy --selector run=envoy --port=10000 --type=NodePort
```


```shell
k exec -it envoy-58bcfd8574-s5gjv -- /bin/bash
ps -ef|grep envoy
# envoy          1       0  0 08:44 ?        00:00:00 envoy -c /etc/envoy/envoy.yaml
```



## 非级联的方式逐个删除对象
```shell
k delete deployment envoy --cascade=orphan
k delete service envoy --cascade=orphan
k delete rs envoy-58bcfd8574 --cascade=orphan

k get pods
k get replicasets
k get endpoints

```

