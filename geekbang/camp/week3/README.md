# goroutine
启动一个goroutine时:
1.它什么时候结束 
2.你有什么办法结束它

## keep yourself busy or do the work yourself
当主goroutine需要等另子goroutine工作完成,等待它结果返回,才能继续执行(一般需要用waitGroup或者channel来实现)
在这种情况下直接在当前goroutine完成这个工作是最简单的,因为委派出去,需要等待另外的goroutine完成,再回调回来

## Leave concurrency to the caller 
让调用者来决定要不要并发,让使用权交给使用者
1. v1:func ListDirectory(dir string) ([]string,error) 返回全部
2. v2:func ListDirectory(dir string) (chan string) 返回channel string,内部维护一个子goroutine
3. v3:func ListDirectory(dir string,fn func(string)) 返回callback,类似filepath.WalkDir()

# dataRace
Detecting Race Conditions With Go

# sync.atomic
典型的场景copy on write
比锁Mutex,RWMutex快,也要根据场景使用

# Mutex
1. Barging. 这种模式是为了提高吞吐量，当锁被释放时，它会唤醒第一个等待者，然后把锁给第一
个等待者或者给第一个请求锁的人。
2. Handsoff. 当锁释放时候，锁会一直持有直到第一个等待者准备好获取锁。它降低了吞吐量，因
为锁被持有，即使另一个 goroutine 准备获取它。
3. Spinning. 自旋在等待队列为空或者应用程序重度使用锁时效果不错。parking 和 unparking 
goroutines 有不低的性能成本开销，相比自旋来说要慢得多。

Go 1.8 使用了 Barging 和 Spining 的结合实现。当试图获取已经被持有的锁时，如果本地队列为空
并且 P 的数量大于1，goroutine 将自旋几次（用一个 P 旋转会阻塞程序）。自旋后，goroutine 
park。在程序高频使用锁的情况下，它充当了一个快速路径。
Go 1.9 通过添加一个新的饥饿模式来解决先前解释的问题，该模式将会在释放时候触发 handsoff。
所有等待锁超过一毫秒的 goroutine（也称为有界等待）将被诊断为饥饿。当被标记为饥饿状态时，
unlock 方法会 handsoff 把锁直接扔给第一个等待者。
在饥饿模式下，自旋也被停用，因为传入的 goroutines 将没有机会获取为下一个等待者保留的锁。

# errGroup
我们把一个复杂的任务，尤其是依赖多个微服务 RPC 需要聚合数据的任务，分解为依赖和并行，
依赖的意思为: 需要上游 a 的数据才能访问下游 b 的数据进行组合。但是并行的意思为: 分解为多
个小任务并行执行，最终等全部执行完毕。
https://pkg.go.dev/golang.org/x/sync/errgroup
核心原理: 利用 sync.Waitgroup 管理并行执行的 goroutine
• 并行工作流
• 错误处理 或者 优雅降级
• context 传播和取消
• 利用局部变量+闭包

errgroup 内部用了WaitGroup来管理goroutine,用sync.Once来做error的处理

# sync.Pool
sync.Pool 的场景是用来保存和复用临时对象，以减少内存分配，降低 GC 压力（Request-Driven 特别合适）。

Get 返回 Pool 中的任意一个对象。如果 Pool 为空，则调用 New 返回一个新创建的对象。

放进 Pool 中的对象，会在说不准什么时候被回收掉。所以如果事先 Put 进去 100 个对象，下次 Get 的时候发现Pool 是空也是有可能的。不过这个特性的一个好处就在于不用担心 Pool 会一直增长，因为 Go 已经帮你在 Pool 中做了回收机制。

这个清理过程是在每次垃圾回收之前做的。之前每次GC 时都会清空 pool，而在1.13版本中引入了 victim cache，会将 pool 内数据拷贝一份，避免 GC 将其清空，即使没有引用的内容也可以保留最多两轮 GC。

# channels

channels 是一种类型安全的消息队列，充当两个goroutine 之间的管道，将通过它同步地进行任意资
源的交换。chan 控制 goroutines 交互的能力从而创建了 Go 同步机制。当创建的 chan 没有容量时，称为无缓冲通道。反过来，使用容量创建的 chan 称为缓冲通道。

要了解通过 chan 交互的 goroutine 的同步行为是什么，我们需要知道通道的类型和状态。无缓冲通道和缓冲通道的使用场景不同，所以让我们单独讨论每个场景。

## Unbuffered Channels

ch := make(chan struct{}) 无缓冲信道的本质是保证同步。

无缓冲 chan 没有容量，因此进行任何交换前需要两个goroutine 同时准备好。当 goroutine 试图将一个资源发送到一个无缓冲的通道并且没有 goroutine 等待接收该资源时，该通道将锁住发送 goroutine 并使其等待。当 goroutine 尝试从无缓冲通道接收，并且没有
goroutine 等待发送资源时，该通道将锁住接收goroutine 并使其等待。

• Receive 先于 Send 发生。
• 好处: 100% 保证能收到。
• 代价: 延迟时间未知。

## Buffered Channels

buffered channel 具有容量，因此其行为可能有点不同。当 goroutine 试图将资源发送到缓冲通道，而该通道已满时，该通道将锁住 goroutine 并使其等待缓冲区可用。如果通道中有空间，发送可以立即进行，goroutine 可以继续。当 goroutine 试图从缓冲通道接收数据，而缓冲通道为空时，该通道将锁住 goroutine 并使其等待资源被发送。

• Send 先于 Receive 发生。
• 好处: 延迟更小。
• 代价: 不保证数据到达，越大的 buffer，越小的保障到达。buffer = 1 时，给你延迟一个消息的保障。

## Go Concurrency Patterns
• Timing out
• Moving on
• Pipeline
• Fan-out, Fan-in
• Cancellation

# Request-scoped context

Go 1.7 引入一个 context 包，它使得跨 API 边界的goroutine 能够很容易享受到请求级别的元数据、取消信号和截止日期传递给处理请求所涉及到的所有goroutine 的功能（显式传递）。

目前有两种方法可以将 context 对象集成到 API 中：
• The first parameter of a function call
首参数传递 context 对象，比如，参考 net 包Dialer.DialContext。此函数执行正常的 Dial 操作，但可以通过context 对象取消函数调用。
• Optional config on a request structure
在第一个 request 对象中携带一个可选的 context 对象。例如net/http 库的 Request.WithContext，通过携带给定的 context 对象，返回一个新的 Request 对象。
(第一种用的比较多,第二种在基础库的设计里面用的比较多)

使用 context 的一个很好的心智模型是让 context 在程序中流动，并贯穿你的所有代码。这通常意味着你不希望将其存储在结构体之中。它从一个函数传递到另一个函数，并根据需要进行扩展。理想情况下，每个请求都会创建一个 context 对象，并在请求结束时过期。
不存储上下文的一个例外是，当你需要将它放入一个结构中时，该结构纯粹用作通过通道传递的消息。