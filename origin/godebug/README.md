# GODEBUG 查看调度跟踪

让 Go 更强大的原因之一莫过于它的 GODEBUG 工具，GODEBUG 的设置可以让 Go 程序在运行时输出调试信息，
可以根据你的要求很直观的看到你想要的调度器或垃圾回收等详细信息，并且还不需要加装其它的插件，非常方便

## schedtrace
```
schedtrace：设置 schedtrace=X 参数可以使运行时在每 X 毫秒发出一行调度器的摘要信息到标准 err 输出中


$ GODEBUG=schedtrace=1000 go run main.go 
SCHED 0ms: gomaxprocs=4 idleprocs=1 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0]
SCHED 1000ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [1 2 2 1]
SCHED 2000ms: gomaxprocs=4 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 [1 2 2 1]

1. sched：每一行都代表调度器的调试信息，后面提示的毫秒数表示启动到现在的运行时间，输出的时间间隔受 schedtrace 的值影响。
2. gomaxprocs：当前的 CPU 核心数（GOMAXPROCS 的当前值）。
3. idleprocs：空闲的处理器数量，后面的数字表示当前的空闲数量。
4. threads：OS 线程数量，后面的数字表示当前正在运行的线程数量。
5. spinningthreads：自旋状态的 OS 线程数量。
6. idlethreads：空闲的线程数量。
7. runqueue：全局队列中中的 Goroutine 数量，而后面的 [0 0 1 1] 则分别代表这 4 个 P 的本地队列正在运行的 Goroutine 数量。
```

## scheddetail
```
scheddetail：设置 schedtrace=X 和 scheddetail=1 可以使运行时在每 X 毫秒发出一次详细的多行信息，信息内容主要包括调度程序、处理器、OS 线程 和 Goroutine 的状态
如果我们想要更详细的看到调度器的完整信息时，我们可以增加 scheddetail 参数，就能够更进一步的查看调度的细节逻辑，如下：

$ GODEBUG=scheddetail=1,schedtrace=1000 go run main.go

P0: status=1 schedtick=1 syscalltick=0 m=0 runqsize=9 gfreecnt=0 timerslen=0
P1: status=0 schedtick=2 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
...
M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=true blocked=true lockedg=-1
...
G1: status=4(semacquire) m=-1 lockedm=-1
G2: status=4(force gc (idle)) m=-1 lockedm=-1
G3: status=4(GC sweep wait) m=-1 lockedm=-1
G4: status=4(GC scavenge wait) m=-1 lockedm=-
...
SCHED 1006ms: gomaxprocs=16 idleprocs=16 threads=19 spinningthreads=0 idlethreads=12 runqueue=0 gcwaiting=0 nmidlelocked=1 stopwait=0 sysmonwait=0

```

# GODEBUG 查看垃圾回收（GC）
```
$ GODEBUG=gctrace=1 go run main.go  

gc 1 @0.004s 2%: 0.063+0.61+0.035 ms clock, 1.0+0.10/0.39/0+0.57 ms cpu, 4->5->1 MB, 5 MB goal, 16 P
gc 2 @0.006s 4%: 0.071+0.55+0.082 ms clock, 1.1+0.087/0.45/0+1.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 3 @0.008s 4%: 0.056+1.1+0.083 ms clock, 0.91+0.16/0.41/0+1.3 ms cpu, 4->5->2 MB, 5 MB goal, 16 P
...
# command-line-arguments
SCHED 0ms: gomaxprocs=16 idleprocs=13 threads=5 spinningthreads=1 idlethreads=1 runqueue=0 [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
# command-line-arguments
gc 1 @0.001s 3%: 0.006+0.33+0.018 ms clock, 0.098+0.049/0.56/0.17+0.29 ms cpu, 4->6->5 MB, 5 MB goal, 16 P
gc 2 @0.006s 1%: 0.006+0.62+0.021 ms clock, 0.10+0.061/0.80/0.16+0.33 ms cpu, 9->10->7 MB, 10 MB goal, 16 P

gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
1. gc#：GC 执行次数的编号，每次叠加。
2. @#s：自程序启动后到当前的具体秒数。
3. #%：自程序启动以来在 GC 中花费的时间百分比。
4. #+...+#：GC 的标记工作共使用的 CPU 时间占总 CPU 时间的百分比。
5. #->#-># MB：分别表示 GC 启动时, GC 结束时, GC 活动时的堆大小.
6. #MB goal：下一次触发 GC 的内存占用阈值。
7. #P：当前使用的处理器 P 的数量。
```