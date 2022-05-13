package main

type Student struct {
	Class int
}

//Demo1_1和Demo1_2 谁执行效率高?

//go tool compile 将源代码编译成 .o 目标文件，并输出汇编代码
//方法一:
// go tool compile -S demo1_1.go demo1.go
// go tool compile -S demo1_2.go demo1.go
//demo1_1.o 和 demo1_2.o 代码完全一直

//方法二: 反汇编，即从可执行文件反编译成汇编，所以要先用 go build 命令编译出可执行文件。
//go build main.go && go tool objdump ./main
