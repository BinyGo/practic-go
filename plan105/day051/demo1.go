package day051

import "fmt"

// /1.下面的代码能否正确输出？
func Demo1() {
	var fn1 = func() {}
	var fn2 = func() {}
	// if fn1 != fn2 { //cannot compare fn1 != fn2 (operator != not defined for func())
	// 	println("fn1 not equal fn2")
	// }
	// 参考答案及解析：编译错误 函数只能与 nil 比较。
	println(fn1, fn2)
}

// /1.下面的代码能否正确输出？
type T struct {
	n int
}

func Demo2() {
	m := make(map[int]T)
	//m[0].n =1 //cannot assign to struct field m[0].n in map
	// map[key]struct 中 struct 是不可寻址的，所以无法直接赋值。
	t := T{1}
	m[0] = t
	fmt.Println(m[0].n) //1

}
