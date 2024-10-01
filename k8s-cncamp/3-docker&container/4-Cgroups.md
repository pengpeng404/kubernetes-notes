# CPU子系统

`默认多核心系统`

在 Cgroup 中有很多的控制组，
CPU子系统只是一部分，
用来控制进程组的CPU资源分配。
进程组内的进程共享资源。

CPU的使用量通常以百分比表示，
用来描述某个进程在一段时间内使用了多少处理器的计算能力。
例如，
如果某个进程占用了100%的CPU，
这意味着该进程在一个核心的时间片内占用了全部可用的计算资源。

## cpu.shares

相对权重控制进程组的CPU资源分配，
用来相对的调节多个进程组之间的资源征用。
默认是1024。

如果一个进程组的 `cpu.shares` 设置为 1024，
而另一个进程组的 `cpu.shares` 设置为 512，
当两个进程组都需要CPU时，
前者会比后者多获得大约两倍的CPU时间。
但这并不是绝对的分配，
它会根据其他进程是否有空闲需求来动态调整。

`cpu.shares`：用于设置进程组之间的相对优先级，
进程组会与其他进程组共享资源。
它在多进程组竞争资源时生效，
表示相对的资源分配。

## cfs_quota_us & cfs_period_us

`cfs_period_us`：指定CFS调度器的时间周期，
单位是微秒（us）。
例如，默认值通常是 100,000us，即100毫秒。

`cfs_quota_us`：指定在一个 `cfs_period_us` 时间周期内，
进程组最多可以使用的CPU时间，
也以微秒为单位。

如果你希望进程组使用多个核心的全部时间，
可以将 `cfs_quota_us` 设置为多个 `cfs_period_us` 的倍数。
例如，设置为 200000 微秒（即2倍周期），
意味着进程组可以使用两个核心的全部CPU时间。

---

`cfs_quota_us` 和 `cfs_period_us`：用于严格控制进程组的CPU使用，
和其他进程组无关。这是硬性限制，
适用于你想要精确控制某个进程组的资源使用情况时。

`cpu.shares`：用于设置进程组之间的相对优先级，
进程组会与其他进程组共享资源。它在多进程组竞争资源时生效，
表示相对的资源分配。

---

## CFS调度器

https://arthurchiao.art/blog/linux-cfs-design-and-implementation-zh/



## CPU子系统练习

linux系统并不会自动创建控制组

```shell
# step 1 在cgroup cpu子系统目录中创建目录结构
cd /sys/fs/cgroup/cpu
mkdir cpudemo
cd cpudemo
# step 2 写一个进程demo eg 死循环
# step 3 查看pid 并把该进程放入 cpudemo 控制组
echo <pid> > cgroup.procs
# step 4 通过修改 cfs 限制资源
echo 100000 > cpu.cfs_quota_us
# 控制该死循环cpu资源为100%
```



# Memory子系统

`memory.limit_in_bytes`：设置Cgroup下进程的最多能使用的内存，
如果设置为-1，则对该Cgroup的内存使用不做限制。
kubernetes使用这个限制pod内存，如果使用内存超过限制，执行OOM。











