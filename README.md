# My_Docker
my docker by golang

### docker详
docker本质其实是一个特殊的进程，
这个进程特殊在它被Namespace和Cgroup技术做了装饰，
Namespace将该进程与Linux系统进行隔离开来，让该进程处于一个虚拟的沙盒中，
而Cgroup则对该进程做了一系列的资源限制，两者配合模拟出来一个沙盒的环境。

### Namespace
Linux对线程提供了六种隔离机制，分别为：uts、pid、user、mount、network ipc，
它们的作用如下：
uts：用来隔离主机名
pid：用来隔离进程PID号的
user：用来隔离用户的
mount：用来隔离各个进程看到的挂载点试图
network：用来隔离网络
ipc：用来隔离System V IPC和POSIX message queues

### Cgroup概念
Linux Cgroup提供了对一组进程及子进程的资源限制，控制和统计能力，这些资源包括CPU，
内存，存储，网络等。通过Cgroup，可以方便的限制某个进程的资源占用，并且可以实时监控进程
和统计信息。

Cgroup完成资源限制主要通过下面三个组件
* cgroup:是对进程分组管理的一种机制
* subsystem:是一组资源控制的模块
* hierarchy:把一组cgroup串成一个树状结构（可让其实现继承）

