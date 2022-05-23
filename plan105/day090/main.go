package main

import "fmt"

type T int
type T2 []int

func F(t T) {}
func F2(t2 T2)

func Demo1() {
	var q int
	//F(q) //cannot use q (variable of type int) as T value in argument to F
	var q2 []int
	F2(q2)
	fmt.Println(q)
	// 我们将这两道题目放到一块做一个解析，第一题不能通过编译，第二题可以通过编译。
	// 我们知道不同类型的值是不能相互赋值的，即使底层类型一样，所以第一题编译不通过；
	// 对于底层类型相同的变量可以相互赋值还有一个重要的条件，即至少有一个不是有名类型（named type）。

	// 这是 Go 语言规范手册的原文：
	// "x's type V and T have identical underlying types and at least one of V or T is not a named type. "
	// Named Type 有两类：
	// 内置类型，比如 int, int64, float, string, bool 等；
	// 使用关键字 type 声明的类型；
	// Unnamed Type 是基于已有的 Named Type 组合一起的类型，例如：struct{}、[]string、interface{}、map[string]bool 等。
}
