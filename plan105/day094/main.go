package main

import "fmt"

//1.下面这段代码输出什么？请简要说明。

func Demo1() {
	a := 2 ^ 15
	b := 4 ^ 15
	if a > b {
		println("a")
	} else {
		println("b")
	}
	// 参考答案及解析：a。Go 语言里面 ^ 表示按位异或，而不是求幂。
	// 0010 ^ 1111 == 1101   (2^15 == 13)
	// 0100 ^ 1111 == 1011   (4^15 == 11)
}

var nil = new(int)

func Demo2() {
	var p *int
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
	//p is not nil
}
