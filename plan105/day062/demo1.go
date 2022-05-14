package day062

import "fmt"

func Demo1() {
	nil := 123
	fmt.Println(nil) //123
	//var _ map[string]int = nil //cannot use nil (variable of type int) as map[string]int value in variable declaration
	//参考答案及解析：当前作用域中，预定义的 nil 被覆盖，此时 nil 是 int 类型值，不能赋值给 map 类型。
}

func Demo2() {
	var x int8 = -128
	var y = x / -1
	fmt.Println(y)
	//参考答案及解析：-128。因为溢出。

	var x2 int32 = -128
	var y2 = x2 / -1
	fmt.Println(y2) //128
}
