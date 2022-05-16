# Prof 
是用于可视化和分析性能分析数据的工具，PProf 以 profile.proto 读取分析样本的集合，并生成报告以可视化并帮助分析数据（支持文本和图形报告）

## 有哪几种采样方式
1. runtime/pprof：采集程序（非 Server）的指定区块的运行数据进行分析。
2. net/http/pprof：基于 HTTP Server 运行，并且可以采集运行时数据进行分析。
3. go test：通过运行测试用例，并指定所需标识来进行采集。

## 支持什么使用模式
1. Report generation：报告生成。
2. Interactive terminal use：交互式终端使用。
3. Web interface：Web 界面。

## 可以做什么
1. CPU Profiling：CPU 分析，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置。
2. Memory Profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏。
3. Block Profiling：阻塞分析，记录 Goroutine 阻塞等待同步（包括定时器通道）的位置，默认不开启，需要调用 runtime.SetBlockProfileRate 进行设置。
4. Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况，默认不开启，需要调用 runtime.SetMutexProfileFraction 进行设置。
5. Goroutine Profiling： Goroutine 分析，可以对当前应用程序正在运行的 Goroutine 进行堆栈跟踪和分析。这项功能在实际排查中会经常用到，因为很多问题出现时的表象就是 Goroutine 暴增，而这时候我们要做的事情之一就是查看应用程序中的 Goroutine 正在做什么事情，因为什么阻塞了，然后再进行下一步。

## http://127.0.0.1:6060/debug/pprof/
1. allocs：查看过去所有内存分配的样本，访问路径为 $HOST/debug/pprof/allocs。
2. block：查看导致阻塞同步的堆栈跟踪，访问路径为 $HOST/debug/pprof/block。
3. cmdline： 当前程序的命令行的完整调用路径。
4. goroutine：查看当前所有运行的 goroutines 堆栈跟踪，访问路径为 $HOST/debug/pprof/goroutine。
5. heap：查看活动对象的内存分配情况， 访问路径为 $HOST/debug/pprof/heap。
6. mutex：查看导致互斥锁的竞争持有者的堆栈跟踪，访问路径为 $HOST/debug/pprof/mutex。
7. profile： 默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件，访问路径为 $HOST/debug/pprof/profile。
8. threadcreate：查看创建新 OS 线程的堆栈跟踪，访问路径为 $HOST/debug/pprof/threadcreate。


## 可视化界面
```
$ wget http://127.0.0.1:6060/debug/pprof/profile  
默认需要等待 30 秒，执行完毕后可在当前目录下发现采集的文件 profile，针对可视化界面我们有两种方式可进行下一步分析

$ go tool pprof -http=:6001 profile 
该命令将在所指定的端口号运行一个 PProf 的分析用的站点

```

## 通过交互式终端使用
### CPU Profiling
```
HTTP: $ go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
TLS: $ go tool pprof https+insecure://localhost:6060/debug/pprof/profile\?seconds\=60

Fetching profile over HTTP from http://localhost:6060/debug/pprof/profile?seconds=60
Saved profile in /home/biny/pprof/pprof.main.samples.cpu.001.pb.gz
File: main
Type: cpu
Time: May 16, 2022 at 10:31am (CST)
Duration: 60s, Total samples = 350ms ( 0.58%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top 10
Showing nodes accounting for 300ms, 85.71% of 350ms total
Showing top 10 nodes out of 51
      flat  flat%   sum%        cum   cum%
      90ms 25.71% 25.71%       90ms 25.71%  syscall.Syscall
      70ms 20.00% 45.71%       70ms 20.00%  runtime.epollwait
      60ms 17.14% 62.86%       60ms 17.14%  runtime.futex
      20ms  5.71% 68.57%       20ms  5.71%  runtime.write1
      10ms  2.86% 71.43%      100ms 28.57%  os.(*File).Write
      10ms  2.86% 74.29%       10ms  2.86%  runtime.acquirep
      10ms  2.86% 77.14%       10ms  2.86%  runtime.checkTimers
      10ms  2.86% 80.00%      150ms 42.86%  runtime.findrunnable
      10ms  2.86% 82.86%       10ms  2.86%  runtime.lock2
      10ms  2.86% 85.71%       10ms  2.86%  runtime.memmove

flat：函数自身的运行耗时。
flat%：函数自身在 CPU 运行耗时总比例。
sum%：函数自身累积使用 CPU 总比例。
cum：函数自身及其调用函数的运行总耗时。
cum%：函数自身及其调用函数的运行耗时总比例。
Name：函数名。
```
### Heap Profiling
```
$ go tool pprof http://localhost:6060/debug/pprof/heap

Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap
Saved profile in /home/biny/pprof/pprof.main.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
File: main
Type: inuse_space
Time: May 16, 2022 at 10:36am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1536.71kB, 100% of 1536.71kB total
Showing top 10 nodes out of 15
      flat  flat%   sum%        cum   cum%
  512.50kB 33.35% 33.35%   512.50kB 33.35%  runtime.allocm
  512.20kB 33.33% 66.68%   512.20kB 33.33%  runtime.malg
  512.01kB 33.32%   100%   512.01kB 33.32%  main.Add (inline)
         0     0%   100%   512.01kB 33.32%  main.main.func1
         0     0%   100%   512.50kB 33.35%  runtime.mstart
         0     0%   100%   512.50kB 33.35%  runtime.mstart0
         0     0%   100%   512.50kB 33.35%  runtime.mstart1
         0     0%   100%   512.50kB 33.35%  runtime.newm
         0     0%   100%   512.20kB 33.33%  runtime.newproc.func1
         0     0%   100%   512.20kB 33.33%  runtime.newproc1

$ go tool pprof -alloc_objects http://localhost:6060/debug/pprof/heap
inuse_space：分析应用程序的常驻内存占用情况。(默认)
alloc_objects：分析应用程序的内存临时分配情况。
另外还有 inuse_objects 和 alloc_space 类别，分别对应查看每个函数所分别的对象数量和查看分配的内存空间大小
```

