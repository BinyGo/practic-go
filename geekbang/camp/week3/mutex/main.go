package main

import (
	"sync"
	"time"
)

func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex

	//goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	//goroutine 2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		mu.Unlock()
	}
	done <- true
	//1.8版本goroutine1拿锁700多万次,goroutine2只拿10次
	/*
		首先，goroutine1 将获得锁并休眠100ms。当 goroutine2 试图获取锁时，它将被添加到
		一个 FIFO 顺序的锁的队列中，goroutine 将进入等待状态。

		然后，当 goroutine1 完成它的工作时，它将释放锁。此版本将通知队列唤醒 goroutine2。
		goroutine2 将被标记为可运行的，并且正在等待 Go 调度程序在线程上运行。

		然而，当 goroutine2 等待运行时，goroutine1将再次请求锁。
		goroutine2 尝试去获取锁，结果悲剧的发现锁又被人持有了，它自己继续进入到等待模式。
	*/
	//1.9后加入饥饿模式后,goroutine1拿锁57次,goroutine2拿10次,解决公平问题
}

// Mutex
// 我们看看几种 Mutex 锁的实现: • Barging. 这种模式是为了提高吞吐量，当锁被释放时，它会唤醒第一个等待者，然后把锁给第一
// 个等待者或者给第一个请求锁的人。
// Handsoff. 当锁释放时候，锁会一直持有直到第一个等待者准备好获取锁。它降低了吞吐量，因
// 为锁被持有，即使另一个 goroutine 准备获取它。
// 一个互斥锁的 handsoff 会完美地平衡两个 goroutine 之间的锁分配，但是会降低性能，因为它会
// 迫使第一个 goroutine 等待锁。 • Spinning. 自旋在等待队列为空或者应用程序重度使用锁时效果不错。parking 和 unparking
// goroutines 有不低的性能成本开销，相比自旋来说要慢得多。
