package day058

import "fmt"

type T struct {
	x int
	y *int
}

func Demo1() {
	i := 20
	t := T{10, &i}
	p := &t.x
	*p++
	*p++
	t.y = p
	fmt.Println(t.y)  //0xc0001164b0
	fmt.Println(*t.y) //12
	// 运算符优先级。如下规则：递增运算符 ++ 和递减运算符 –- 的优先级低于解引用运算符 * 和取址运算符 &，
	// 解引用运算符和取址运算符的优先级低于选择器 . 中的属性选择操作符。
}

func Demo2() {
	x := make([]int, 2, 10)
	_ = x[6:10]
	_ = x[6:] //runtime error: slice bounds out of range [6:2]
	_ = x[2:]
	//参考答案：第 26 行，截取符号 [i:j]，如果 j 省略，默认是原切片或者数组的长度，x 的长度是 2，小于起始下标 6 ，所以 panic。
}
