package main

import (
	"fmt"
	"log"
	"path"
	"runtime"
)

/*
Go 是一门有 runtime 的语言，什么是 runtime？其实就是一段辅助程序，用户没有写的代码，runtime 替我们写了，比如 Go 调度器的代码。
我们只需要知道用 go 关键字创建 goroutine，就可以疯狂堆业务了。至于 goroutine 是怎么被调度的，根本不需要关心，这些是 runtime 调度器的工作。
那我们自己写的代码如何和 runtime 里的代码对应起来呢？
*/

func main() {
	go func() {
		println(1 + 2)
	}()
	CallerA()
}

// go tool compile 将源代码编译成 .o 目标文件，并输出汇编代码
// go tool compile -S main.go
// CALL    runtime.newproc(SB) //能看到 go func(){}() 对应 newproc() 函数
// 这时再深入研究下 newproc() 函数就大概知道 goroutine 是如何被创建的。
// go tool compile -S main.go | grep "main.go:4"

// 方法二: 反汇编，即从可执行文件反编译成汇编，所以要先用 go build 命令编译出可执行文件。
// go build main.go && go tool objdump ./main

// runtime.Caller 报告当前 Go 程序调用栈所执行的函数的文件和行号信息
func CallerA() {
	//获取的是 CallerA 这个函数的调用栈
	pc, file, lineNo, ok := runtime.Caller(0)
	//获取的是 CallerA函数的调用者的调用栈
	pc1, file1, lineNo1, ok1 := runtime.Caller(1)
	fmt.Println(pc, file, lineNo, ok) //栈标识符 带路径的完整文件名 该调用在文件中的行号 获取成功失败状态bool
	if !ok {
		log.Fatal("runtime.Caller() failed")
		return
	}
	fmt.Println(pc1, file1, lineNo1, ok1)
	if !ok1 {
		log.Fatal("runtime.Caller() failed")
		return
	}

	//runtime.FuncForPC 函数返回一个表示调用栈标识符pc对应的调用栈的*Func；如果该调用栈标识符没有对应的调用栈，函数会返回nil。
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	fmt.Printf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
	//FuncName:github.com/practic-go/compile/runtime.CallerA, file:main.go, line:35
}
