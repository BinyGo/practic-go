package day028

import "fmt"

func Demo1() {
	//go 中不同类型是不能比较的，而数组长度是数组类型的一部分，所以 […]int{1} 和 [2]int{1} 是两种不同的类型，不能比较；
	//fmt.Println([...]int{1}==[2]int{1}) //cannot compare ([...]int literal) == ([2]int literal) (mismatched types [1]int and [2]int)

	//切片是不能比较的；
	//fmt.Println([]int{1} == []int{1}) //cannot compare ([]int literal) == ([]int literal) (operator == not defined for []int)
}

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	fmt.Println(*p)
}

func Demo2() {
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar() //panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(*p)
	// A. 5 5
	// B. runtime error
	// 参考答案及解析：B。知识点：变量作用域。问题出在操作符:=，对于使用:=定义的变量，如果新变量与同名已定义的变量不在同一个作用域中，那么 Go 会新定义这个变量。
	// 对于本例来说，main() 函数里的 p 是新定义的变量，会遮住全局变量 p，导致执行到bar()时程序，全局变量 p 依然还是 nil，程序随即 Crash。
	// 正确的做法是将 main() 函数修改为：
}

//正确的做法修改为：
func Demo3() {
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
