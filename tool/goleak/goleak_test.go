package main

import (
	"testing"

	"go.uber.org/goleak"
)

func Test_leak(t *testing.T) {
	defer goleak.VerifyNone(t)
	leak()
}

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

//除了 uber-go/goleak，也还有 ysmood/gotrace 等同类型的库能够达到类似的效果。
