package main

import "fmt"

//1.下面代码输出什么？
func Demo1() {
	a := 1
	for i := 0; i < 5; i++ {
		a := a + 1
		a = a * 2 //this value of a is never used (SA4006)
	}
	fmt.Println(a) //1
	//参考答案及解析：1。知识点：变量的作用域。注意 for 语句的变量 a 是重新声明，它的作用范围只在 for 语句范围内。
}

//2.下面代码输出什么？

func test(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		// ret := 10 //ret declared but not used
		// 编译错误。知识点：变量的作用域。编译错误信息：ret is shadowed during return。
		ret = 10
		return
	}
	return
}

func Demo2() {
	result := test(10)
	fmt.Println(result)
}
