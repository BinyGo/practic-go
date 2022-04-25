package main

import "fmt"

func Demo1() {
	list := new([]int)
	//list = append(list, 1) //invalid argument: list (variable of type *[]int) is not a slice
	//不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。
	//可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
	fmt.Println(list)

}

func Demo2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//s1 = append(s1, s2) 不能通过编译。append() 的第二个参数不能直接使用 slice，
	//需使用 … 操作符，将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)。
	//s1 = append(s1, s2...)
	fmt.Println(s1, s2)

}

var (
	//size     := 1024 不能通过编译。这道题的主要知识点是变量声明的简短模式，形如：x := 100。但这种声明方式有限制：
	// 1.必须使用显示初始化；
	// 2.不能提供数据类型，编译器会自动推导；
	// 3.只能在函数内部使用简短模式；
	size     = 1024
	max_size = size * 2
)

func Demo3() {
	fmt.Println(size, max_size)
}
