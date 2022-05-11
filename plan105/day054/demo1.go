package day054

import "fmt"

//1.下面的代码有什么问题？
type N int

func (n N) Value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) Pointer() {
	*n++
	fmt.Printf("v:%p,%v\n", n, *n)
}
func Demo1() {
	var a N = 25
	p := &a
	p1 := &p
	//p1.Value()   //p1.value undefined (type **N has no field or method value)
	//p1.Pointer() //p1.pointer undefined (type **N has no field or method pointer)
	//不能使用多级指针调用方法。
	fmt.Println(a, p, p1) //25 0xc0000be070 0xc0000a8020
}

//2.下面的代码输出什么？
func (n N) test() {
	fmt.Println(n)
}
func Demo2() {
	var n N = 10
	fmt.Println(n)

	n++
	f1 := N.test
	f1(n)

	n++
	f2 := (*N).test
	f2(&n)
	// 参考答案及解析：10 11 12。
	// 知识点：方法表达式。通过类型引用的方法表达式会被还原成普通函数样式，接收者是第一个参数，调用时显示传参。
	// 类型可以是 T 或 *T，只要目标方法存在于该类型的方法集中就可以。
}

//还可以直接使用方法表达式调用：
func Demo3() {
	var n N = 10
	fmt.Println(n)

	n++
	N.test(n)

	n++
	(*N).test(&n)
}
