package main

//dlv 调试工具--可以调试 Go、以及和 Go 程序互动的方法

/*
向一个 nil 的 slice append 元素，不会有任何问题。
但是向一个 nil 的 map 插入新元素，马上就会报 panic。这是为什么呢？又是在哪 panic 呢？
首先写出让 map 产生 panic 的示例程序：
*/

func main() {
	var m map[int]int
	m[1] = 1
}

/*
1. build命令编译生成可执行文件
$ go build main.go

2.使用 dlv 进入调试状态
$ dlv exec ./main

3.使用 b 这个命令打断点，有三种方法：
b + 地址
b + 代码行数
b + 函数名
$ b main.go:13 //加断点
$ c //直接运行到断点处
$ disass //执行 disass 命令，可以看到汇编指令
$ si  //这时使用 si 命令，执行单条指令，多次执行 si(这里执行了3次)，就会执行到 map 赋值函数 mapassign_fast64
$ s   //这时再用单步命令 s，就会进入判断 h 的值为 nil 的分支，然后执行 panic 函数
至此，向 nil 的 map 赋值时，产生 panic 的代码就被我们找到了。接着，按图索骥找到对应 runtime 源码的位置，就可以进一步探索了。
*/
/* func mapassign_fast64(t *maptype, h *hmap, key uint64) unsafe.Pointer {
=>  93:         if h == nil {
	94:                 panic(plainError("assignment to entry in nil map"))
	95:         }
*/

//除此之外，我们还可以使用 bt 命令看到调用栈：
//使用 frame 1 命令可以跳转到相应位置。这里 1 对应图中的 a.go:5，也就是我们前面打断点的地方
