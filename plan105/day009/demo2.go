package day009

import "fmt"

//下面这段代码输出什么？

func hello(num ...int) {
	num[0] = 18
}

func Demo2() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
	// A.18
	// B.5
	// C.Compilation error
	// 参考答案及解析：18。知识点：切片,指针,可变函数。
}
