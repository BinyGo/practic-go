# trace 跟踪剖析

有时候单单使用 pprof 还不一定足够完整观查并解决问题，因为在真实的程序中还包含许多的隐藏动作，
例如 Goroutine 在执行时会做哪些操作？执行/阻塞了多长时间？在什么时候阻止？在哪里被阻止的？
谁又锁/解锁了它们？GC 是怎么影响到 Goroutine 的执行的？这些东西用 pprof 是很难分析出来的

# 
```
生成跟踪文件
$ go run main.go 2> trace.out

启动可视化界面
$ go tool trace trace.out

1. View trace：查看跟踪
2. Goroutine analysis：Goroutine 分析
3. Network blocking profile：网络阻塞概况
4. Synchronization blocking profile：同步阻塞概况
5. Syscall blocking profile：系统调用阻塞概况
6. Scheduler latency profile：调度延迟概况
7. User defined tasks：用户自定义任务
8. User defined regions：用户自定义区域
9. Minimum mutator utilization：最低 Mutator 利用率
```
