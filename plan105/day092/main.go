package main

import "fmt"

//2.min() 函数是求两个数之间的较小值，能否在 该函数中添加一行代码将其功能补全。

func min(a int, b uint) {
	var min = 0
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func Demo1() {
	min(1225, 256)
}

//参考答案即解析：利用 copy() 函数的功能：切片复制，并且返回两者长度的较小值。
func min2(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func Demo2() {
	min2(1225, 256)
}
