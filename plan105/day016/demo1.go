package day016

import "fmt"

//1.切片 a、b、c 的长度和容量分别是多少？
func Demo1() {

	//因为它内部实现的复杂性，无法“零值可用”。map 类型变量进行显式初始化后才能使用
	var m map[string]int
	//m["a"] = 1           //assignment to nil map (SA5000) panic: assignment to entry in nil map
	// if v := m["b"]; v != nil { //cannot convert nil (untyped nil value) to int
	// 	fmt.Println(v)
	// }
	if v, ok := m["b"]; ok { //comma ok写法
		fmt.Println(v)
	}
	fmt.Println(m)

	//显式初方法一：使用复合字面值初始化 map 类型变量。
	m1 := map[string]int{}
	m1["a"] = 1
	//显式初方法二：使用 make 为 map 类型变量进行显式初始化。
	m2 := make(map[string]int)
	//m2 := make(map[string]int, 8)  指定初始容量为8
	m2["a"] = 1
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func Demo2() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	//fmt.Println(a.ShowB()) //a.ShowB undefined (type A has no field or method ShowB)
	//fmt.Println(b.ShowA()) //b.ShowA undefined (type B has no field or method ShowA)
	fmt.Println(b.ShowB())
	//A. 23 13
	//B. compilation error
	//参考答案及解析：B。知识点：接口的静态类型。a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；
	//a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错。
}
