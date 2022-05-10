package main

import (
	"fmt"
	"runtime"
	"time"
)

// gorourtine的设计是Go语言并发实现的核心组成部分，易上手，但是也会遭遇各种疑难杂症，其中goroutine泄漏就是重症之一，其出现往往需要排查很久，
// 有人说可以使用pprof来排查，虽然其可以达到目的，但是这些性能分析工具往往是在出现问题后借助其辅助排查使用的，有没有一款可以防患于未然的工具吗？
// 当然有，goleak他来了，其由 Uber 团队开源，可以用来检测goroutine泄漏，并且可以结合单元测试，可以达到防范于未然的目的，本文我们就一起来看一看goleak。

func leak() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
	//<-ch //注释,未取ch,将造成内存泄漏,在单元测试中查看详情
}

// goroutine泄漏
// 不知道你们在日常开发中是否有遇到过goroutine泄漏，goroutine泄漏其实就是goroutine阻塞，
// 这些阻塞的goroutine会一直存活直到进程终结，他们占用的栈内存一直无法释放，从而导致系统的可用内存会越来越少，直至崩溃！简单总结了几种常见的泄漏原因：

// Goroutine内的逻辑进入死循坏，一直占用资源
// Goroutine配合channel/mutex使用时，由于使用不当导致一直被阻塞
// Goroutine内的逻辑长时间等待，导致Goroutine数量暴增
// 接下来我们使用Goroutine+channel的经典组合来展示goroutine泄漏；
func GetData() {
	var ch chan struct{}
	go func() {
		<-ch
	}()
}

func Demo1() {
	defer func() {
		fmt.Println("groutines: ", runtime.NumGoroutine())
	}()
	GetData()
	time.Sleep(time.Second * 2)
}
