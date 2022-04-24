package day001

import "fmt"

// 1.使用值为 nil 的 slice、map会发生啥
// 允许对值为 nil 的 slice 添加元素，但对值为 nil 的 map 添加元素，则会造成运行时 panic。

func ErrDemo1() {
	var m map[string]int
	m["one"] = 1
	// error: panic: assignment to entry in nil map
	// m := make(map[string]int)// map 的正确声明，分配了实际的内存
	fmt.Println(m)
}

// slice 正确示例
func Demo1() {
	var s []int
	s = append(s, 1)
	fmt.Println(s)
}
