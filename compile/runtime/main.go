package main

/*
Go 是一门有 runtime 的语言，什么是 runtime？其实就是一段辅助程序，用户没有写的代码，runtime 替我们写了，比如 Go 调度器的代码。
我们只需要知道用 go 关键字创建 goroutine，就可以疯狂堆业务了。至于 goroutine 是怎么被调度的，根本不需要关心，这些是 runtime 调度器的工作。
那我们自己写的代码如何和 runtime 里的代码对应起来呢？
*/

func main() {
	go func() {
		println(1 + 2)
	}()
}

// go tool compile 将源代码编译成 .o 目标文件，并输出汇编代码
// go tool compile -S main.go
// CALL    runtime.newproc(SB) //能看到 go func(){}() 对应 newproc() 函数
// 这时再深入研究下 newproc() 函数就大概知道 goroutine 是如何被创建的。
// go tool compile -S main.go | grep "main.go:4"

// 方法二: 反汇编，即从可执行文件反编译成汇编，所以要先用 go build 命令编译出可执行文件。
// go build main.go && go tool objdump ./main
