# Kubernetes 面试问题

参考脉脉用户问题


#### 1、docker 和虚拟机有哪些不同
```log
Docker 是轻量级的沙盒，
在其中运行的只是应用，
而虚拟机里面还需要额外的运行独立的操作系统，
非常消耗内存

```

#### 2、Kubernetes 和 docker 关系





#### 3、ipvs 和 iptables 区别
```log
iptables:
- 优点
  - 灵活 功能强大 （tcp不同阶段对包进行操作）
- 缺点
  - 表中的规则过多 响应变慢 线性延时
ipvs:
- 优点
  - 转发效率高 调度算法丰富
- 缺点
  - 对内核版本有要求
```

#### 4、ipvs 和 iptables 切换


#### 5、微服务部署中的蓝绿发布
```log
蓝绿部署中，一共有两套系统
正在服务的系统 -- 绿色
准备发布的系统 -- 蓝色
需要修改都在新版本的蓝色系统上修改，
蓝色系统经过测试，修改，验证，
达到上线标准后，直接把流量切换到蓝色系统

特点：无需停机，风险最小

```

#### 6、蓝绿发布优势与不足



















































