package main

import (
	"testing"

	"go.uber.org/goleak"
)

func Test_leak(t *testing.T) {
	defer goleak.VerifyNone(t)
	leak()
	/*
	   单纯 leak() 正常通过测试
	   加入 defer goleak.VerifyNone(t)
	   可以从报告中看到，运行结构会明确的告诉你发现泄露的 goroutine 的代码堆栈和泄露类型，非常的省心。
	   --- FAIL: Test_leak (0.44s)
	       leaks.go:78: found unexpected goroutines:
	           [Goroutine 19 in state chan send, with github.com/practic-go/tool/goleak.leak.func1 on top of the stack:
	           goroutine 19 [chan send]:
	           github.com/practic-go/tool/goleak.leak.func1()
	                   /golang/src/github.com/practic-go/tool/goleak/goleak.go:6 +0x25
	           created by github.com/practic-go/tool/goleak.leak
	                   /golang/src/github.com/practic-go/tool/goleak/goleak.go:5 +0x6f
	           ]
	   FAIL
	   exit status 1
	   FAIL    github.com/practic-go/tool/goleak       0.445s
	*/
}

func Test_GetData(t *testing.T) {
	//使用goleak主要关注两个方法即可：VerifyNone、VerifyTestMain，VerifyNone用于单一测试用例中测试，VerifyTestMain可以在TestMain中添加，可以减少对测试代码的入侵，举例如下：
	//使用VerifyNone:
	defer goleak.VerifyNone(t)
	GetData()
	//运行结果:
	/*         leaks.go:78: found unexpected goroutines:
	           [Goroutine 19 in state chan send, with github.com/practic-go/tool/goleak.leak.func1 on top of the stack:
	           goroutine 19 [chan send]:
	           github.com/practic-go/tool/goleak.leak.func1()
	                   /golang/src/github.com/practic-go/tool/goleak/goleak.go:16 +0x25
	           created by github.com/practic-go/tool/goleak.leak
	                   /golang/src/github.com/practic-go/tool/goleak/goleak.go:15 +0x6f

	            Goroutine 6 in state chan receive (nil chan), with github.com/practic-go/tool/goleak.GetData.func1 on top of the stack:
	           goroutine 6 [chan receive (nil chan)]:
	           github.com/practic-go/tool/goleak.GetData.func1()
	                   /golang/src/github.com/practic-go/tool/goleak/goleak.go:32 +0x1f
	           created by github.com/practic-go/tool/goleak.GetData
	                   /golang/src/github.com/practic-go/tool/goleak/goleak.go:31 +0x45
	           ]
	   FAIL */
}

// 总结一下goleak的实现原理：
// 使用runtime.Stack()方法获取当前运行的所有goroutine的栈信息，
// 默认定义不需要检测的过滤项，默认定义检测次数+检测间隔，不断周期进行检测，
// 最终在多次检查后仍没有找到剩下的goroutine则判断没有发生goroutine泄漏。
// 资料原文:https://mp.weixin.qq.com/s/GQFuVTurmY0m8WByB1nKeQ

//除了 uber-go/goleak，也还有 ysmood/gotrace 等同类型的库能够达到类似的效果。
