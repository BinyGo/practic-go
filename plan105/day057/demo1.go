package day057

import "fmt"

//1.下面哪一行代码会 panic，请说明原因？

func Demo1() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x //identical expressions on the left and right side of the '==' operator (SA4000)
	_ = x == y
	//_ = y == y
	// 运行报错 panic: runtime error: comparing uncomparable type []int
	// 编译仅警告:identical expressions on the left and right side of the '==' operator (SA4000)
}

var o = fmt.Print

func Demo2() {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c:
			o(2)
			c = nil
		case c <- 1:
			o(3)
		}
	}
	// 321
	// 第一次循环，写操作已经准备好，执行 o(3)，输出 3；
	// 第二次，读操作准备好，执行 o(2)，输出 2 并将 c 赋值为 nil；
	// 第三次，由于 c 为 nil，走的是 default 分支，输出 1。
}
