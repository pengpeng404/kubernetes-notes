# kubectl 基本使用


```shell
# -v 9 打开 debug log
k get ns default -v 9
# 读配置文件
# 发请求 
```
```log
root@master:/home/cadmin# k get ns default -v 9
I1004 07:28:46.671952  790623 loader.go:395] Config loaded from file:  /root/.kube/config
I1004 07:28:46.674471  790623 round_trippers.go:466] curl -v -XGET  
    -H "Accept: application/json;as=Table;v=v1;g=meta.k8s.io,application/json;as=Table;v=v1beta1;g=meta.k8s.io,application/json" 
    -H "User-Agent: kubectl/v1.30.5 (linux/amd64) kubernetes/74e84a9" 'https://192.168.34.101:6443/api/v1/namespaces/default'
I1004 07:28:46.674766  790623 round_trippers.go:510] HTTP Trace: Dial to tcp:192.168.34.101:6443 succeed
I1004 07:28:46.679091  790623 round_trippers.go:553] GET https://192.168.34.101:6443/api/v1/namespaces/default 200 OK in 4 milliseconds
I1004 07:28:46.679494  790623 round_trippers.go:570] HTTP Statistics: DNSLookup 0 ms Dial 0 ms TLSHandshake 2 ms ServerProcessing 1 ms Duration 4 ms
I1004 07:28:46.679755  790623 round_trippers.go:577] Response Headers:
I1004 07:28:46.679997  790623 round_trippers.go:580]     Audit-Id: 21c07b0a-32bf-4217-acfb-c24d58a83aac
I1004 07:28:46.680239  790623 round_trippers.go:580]     Cache-Control: no-cache, private
I1004 07:28:46.680536  790623 round_trippers.go:580]     Content-Type: application/json
I1004 07:28:46.680776  790623 round_trippers.go:580]     X-Kubernetes-Pf-Flowschema-Uid: 3dbda640-59c0-46c1-a5cd-48e328cc2e0c
I1004 07:28:46.681014  790623 round_trippers.go:580]     X-Kubernetes-Pf-Prioritylevel-Uid: aae461c1-438a-4d20-8322-3308f9df6990
I1004 07:28:46.681249  790623 round_trippers.go:580]     Content-Length: 1679
I1004 07:28:46.681482  790623 round_trippers.go:580]     Date: Fri, 04 Oct 2024 07:28:46 GMT
I1004 07:28:46.681744  790623 request.go:1212] Response Body: {
    "kind":"Table",
    "apiVersion":"meta.k8s.io/v1",
    "metadata":{"resourceVersion":"33"},
    "columnDefinitions":[{"name":"Name","type":"string","format":"name","description":"Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names","priority":0},{"name":"Status","type":"string","format":"","description":"The status of the namespace","priority":0},{"name":"Age","type":"string","format":"","description":"CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata","priority":0}],"rows":[{"cells":["default","Active","13d"],"object":{"kind":"PartialObjectMetadata","apiVersion":"meta.k8s.io/v1","metadata":{"name":"default","uid":"6ec49143-7f1e-45a4-a6c5-8444a04f6f4e","resourceVersion":"33","creationTimestamp":"2024-09-21T02:49:55Z","labels":{"kubernetes.io/metadata.name":"default"},"managedFields":[{"manager":"kube-apiserver","operation":"Update","apiVersion":"v1","time":"2024-09-21T02:49:55Z","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:labels":{".":{},"f:kubernetes.io/metadata.name":{}}}}}]}}}]}
NAME      STATUS   AGE
default   Active   13d
```

```shell
cat /root/.kube/config
```
```log
root@master:/home/cadmin# cat /root/.kube/config
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCVENDQWUyZ0F3SUJBZ0lJVXNZdzYrVkxpZmN3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TkRBNU1qRXdNalF6TUROYUZ3MHpOREE1TVRrd01qUTRNRE5hTUJVeApFekFSQmdOVkJBTVRDbXQxWW1WeWJtVjBaWE13Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLCkFvSUJBUUNtN0ZZWGNtSXEzSk02WGtBMCtBRW5xNWQrV0djVDNlSmRRTUp3eUV3K2pTNnlyUG9YbE1aam9qemQKT1N2WW9Rc3g1dXRPd0VlR0VuRGZpejU2Um1FUDZKSWd0S0k4TUswQnUySTYrK2JRMHJDdVRIZ251QjBsclBncwpMQVpORE1VVG9kK1F2REFGUTVYMTB3UzVpalE5VTdMdmJpQm12a3krZ3p0OUY5a2UzOCtrUzR5NDJ6SzIzYzV2CmpkT09LS25CQXdzK0RQakVLaDZSUnVBQWp2cTlOa3ZOODliWW1oRW1NTitiamxEVjQxME9FNzlqOVdqRE1ScXUKSU02SWZEa0pGRWlYR0lNcUlxQnhEazg5cngzVnowWTZtTFhZdXVKa2tpeEt1SGtFb2Z0NVFYWDJna3FqU0JCSgpBdlNXVk10VHBQSi9zcnUrN2gyY3UvTWdLd0FmQWdNQkFBR2pXVEJYTUE0R0ExVWREd0VCL3dRRUF3SUNwREFQCkJnTlZIUk1CQWY4RUJUQURBUUgvTUIwR0ExVWREZ1FXQkJSdE9Zc3IyeFZrWGJEREdtVnl3OVFuUHhuQXVUQVYKQmdOVkhSRUVEakFNZ2dwcmRXSmxjbTVsZEdWek1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQjdFMmFCSmFJYQpUejl5bEtmbTVGUjFhdVVidThHczJKUkZkaDlkaHRIUWFRek1XZFFzVzd1eXArYlM5eGVPRFBPY3FFNG9UZzhTCjdXNVVxSE9LZWtRR3Q3Q0FtZHVQYndMTTRkYkRaSVBTakgrSXVXTWxDc2lWNDBtbzY5cmZhaE1SNDVBSVB3blMKYnB6MzN5SFJHbXZNUHphYkEzYnUrMlB4T1BFWitlVDdHZEVuZWdKWW10ZGp6N0NFdEhYcDZZN044RmpaNVI2dApMZmEveUNWSm50S0FaSitpU2ZEWkZ6Rm5ONmhCTTU2VHRqOGF3c1dxSTZMYUFvYnZWSnlJQjhXcFFISFZWZXNCCk5nUjZROFZPR3VEQTk5bnczdU5QR2ZzcVJvMzV0VkZzTFBNdlhWUTg3NE51cU93UFQ3R2RMcE9zdGVNcUlRVEwKQ1pMeU0zeWpKMWozCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://192.168.34.101:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURLVENDQWhHZ0F3SUJBZ0lJZVY0bEQxc0NabU13RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TkRBNU1qRXdNalF6TUROYUZ3MHlOVEE1TWpFd01qUTRNRFJhTUR3eApIekFkQmdOVkJBb1RGbXQxWW1WaFpHMDZZMngxYzNSbGNpMWhaRzFwYm5NeEdUQVhCZ05WQkFNVEVHdDFZbVZ5CmJtVjBaWE10WVdSdGFXNHdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDN0RPMFgKMlBmb0VZeUN1TkdvQmdxVW95Y3hDc0FTYkkvNHJtT3VBN3RWTHMyN1pVcWdpd0Q2TmhXZUVrVGYxc1ZYWHFaNgpJdkhSTkV6K0JlQS83OGhKWGFEc2JDa3pBVW1OTE4yVlJLZlBZVTZ6cTI0RndsSG1DbjJIbnduYXJOOXFBa3NSCnNMQ0M0OHNQVUltbDRYVFhJcUNQZmFORVZJaFdmU3pvSUNZY2w2RzFHVk05Mjc3aUl6bVd6dzRSMjRKQlZwbDcKcWlJajJORmw0emhPN3NSd0ZjWkdEb003T2pRTksvMTRkcWVZR3NvVkJrRmRETjFFL0lkTUxEQ2o3ZU4rK242dQpvcU1CZmY3WkFwUUpXRVBkRWlLUlluWkI4dC8wV2x5UFR2MWs0Y2swVk9sbTlERlF4S0pDb0UxMHEvOHF0OTYvClNNRnU2SzNHeHR2UFRYdnZBZ01CQUFHalZqQlVNQTRHQTFVZER3RUIvd1FFQXdJRm9EQVRCZ05WSFNVRUREQUsKQmdnckJnRUZCUWNEQWpBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRkcwNWl5dmJGV1Jkc01NYQpaWExEMUNjL0djQzVNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUFxRVBZNURLU1dnNHc0V3N1R3o3YndWQTZPCmE5a0hrK3RiK3lvaHVUR2NLTVVrVDFjc0orejhhYUVMWlhUZm9UK1pkemcxMzNHbGNURURmbGNheFh0dmlaWWIKM093b1ZZMDN6Zmh3dkplUXR3d3UxR2RrRUdQT1ltUU1sVnRnaHpqN1ZFbVB4NGVnbW5lTHpoZEdvUUZoUmc4UwpTQ2toSmFGanBDcVUwWmp4dnRFNVRyYmNQMFN5LzZBYmNWQWgrMEVndXNVS2hTbGpObFJQUjIzNS9KUER4UmJlCk9rSGpESi8yeHZmck1WaEQ2RnZqTkdNS0llTFR1b1F4V0N1UW90VGRNMU5PYW9LeHlBT0VHckpvRXZJaExzYWMKNzJtcll1Yk9oR0pOZkcweFdHWEIvYklEOXNXKzdnR2l5QWJnWGtxNGhyR1pxTGF1aW5JdzJNOXR4WThjCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdXd6dEY5ajM2QkdNZ3JqUnFBWUtsS01uTVFyQUVteVArSzVqcmdPN1ZTN051MlZLCm9Jc0ErallWbmhKRTM5YkZWMTZtZWlMeDBUUk0vZ1hnUCsvSVNWMmc3R3dwTXdGSmpTemRsVVNuejJGT3M2dHUKQmNKUjVncDloNThKMnF6ZmFnSkxFYkN3Z3VQTEQxQ0pwZUYwMXlLZ2ozMmpSRlNJVm4wczZDQW1ISmVodFJsVApQZHUrNGlNNWxzOE9FZHVDUVZhWmU2b2lJOWpSWmVNNFR1N0VjQlhHUmc2RE96bzBEU3Y5ZUhhbm1CcktGUVpCClhRemRSUHlIVEN3d28rM2pmdnArcnFLakFYMysyUUtVQ1ZoRDNSSWlrV0oyUWZMZjlGcGNqMDc5Wk9ISk5GVHAKWnZReFVNU2lRcUJOZEt2L0tyZmV2MGpCYnVpdHhzYmJ6MDE3N3dJREFRQUJBb0lCQUJVSGR3QmZYRCt5K1dFWQo5amsxdEtyUlRPNnVqcm1EaXd3aVR3S1pXTTVTM0w1Y3ZPOGZzWlJ2MEM1ZFQzRDY3R1RPTjFrejVJdm9uVjlSCnVjeDJZTVlleUtETDZEWGJ4ekVnQWlsdDlvL1NHTThLVHV4RzFINVFYNXlIdk12Zzg1MHZkTkVnVkRmaTlFbGMKZkoweG83a3NJM0QrWndTZm5GUmM3bGVLaGI1Zy9FUGh0V2YzMjNEMkk1V1JHNVdxUFdxbXRzNTVqNUFic3ZLbwpVZG9jNFhuUXhOaGF1dU1IWVRrVGIzRVRhcG5ZODdZSDV4NTVNWmp3Y1RoU0ozSHRZelpYSlJWRXA1Z3JENkUwCmRidnE2S3pDZmlZRlU1QjA2N2hHRXBmZEhMN0hUQXVOL2VJdGFvZkxMelFlckdTZmdKOTVyM21WMDQ0QmZFTm0KTFIrQnFBRUNnWUVBdytzd0UyN3BCNG5Fc1V6c3pTdmJWR2FFbjVCWGwyWExwU055eiswYzB0T01iVEpKdzdmQgpYREtmT0t0czFLQWp6LytwMS82SlBmY0FwQS9nTEU1dE5MbU83MmRVQ3E1VWVNZGdGNEVWdkZFM2t2TnlqMDdiCnR4d3RIaXhJeENDOG1GSXhvQzVHaXFZUDVVMlgvNkxnSnhDd2xDMjl2WFNTMTAySDI3YTY1WUVDZ1lFQTlHbUgKcEJmZVRvYTd4TUUrTjhTNG92dkErRmxobnJLS3REMExXaDhuYmk2cU1aRnJzRTIvVEFaT05jWDlkMVpZNXplegp5aWUrMHpWejVBcjdJNW45bGVjL3pSd3B2RXE3TGwwWTd3bkVMVGlpVC9KUTJzdmhwWGVwcTJzWWFLbDFzMnpFCnV1YVEyNU52MXRUYzdMS0hqOTR2Nk9Rd3h6L3Roa0IxTHliS2VXOENnWUVBd2dDaUdwdzVKTDNIaDhva3N3WTYKcWRqYWV5YnpsWGUzc0U3cDRmdHFEMXBzTTdVWVZqWWZ3cXhkL2ovQ0JNcU9xK2oreG1QR2d4V1VET0dybWpRTwo3NmJQWTBGdWR5VXBnRy90TjFrYnJONi9xVVJvckgvcUVlaFV4UXdWQWlGb24yekV0MWtiZ1MvdmphRElZdHRtCkcwanJrYys1azJGY0J0Yy9NTkpCUUFFQ2dZQlJRbWk3Y01nVGVZNGlDMUdCUHlGWDVyV3duQjd0b1ZTbU9nbDEKTEJoeTlJYlhOZzhFcmNTbEpRK0pwMHJ2QzBGQmxtNXJEcTNPRU41MytnS25Rb0poL1dGajh6SVpEUXVRalpsRgptQXltTUVjZXAyU2thZGFhcWQ4NlE3LzR4Q2FDd2UzaWFkZk5lUVpjK2FaOTk1bEVoczJNODVrWUZiUUZ4NVp1ClY4cEhkd0tCZ0M5RjdISFllMkcwMmt5VDQ4QWZmYnZMRVdPaVJ0U003bG8wQTQ2aFIvM1c0U1RBbWI2QzFPWEEKQjVtT2M1SEhtYW9vNG00UDIzOXpMMTB5YTBwbGV4Ym9RMlpRMjlRMVB5S1VmWlRCQ2hkR29PdnhBcTg4TXI3Lwp0dXR5SE53N245S0lLU3JmMFBhNWZIeHhOdi96c1hiV1c3M2JXanlzUUNSeUNqSmJTNGRCCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
```


- `cluster`：集群列表 可能有多个集群
- `contexts`：用什么用户链接上面的哪个集群
- `current-context`：如果不指定哪个 `context` 就默认这个 `context`
- `users`：用户信息


组装成一个 `REST` 调用发给 `apiserver`，
`apiserver` 的地址、用户信息放在 `kubeconfig`



### 常用命令

```shell
k api-resources
```
```shell
ks exec -it etcd-master -- /bin/sh
```
```shell
k describe
```
```shell
# 看标准输出
k logs
```










































