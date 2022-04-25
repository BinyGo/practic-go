package day006

import "fmt"

// 1.通过指针变量 p 访问其成员变量 name，有哪几种方式？
// A.p.name
// B.(&p).name
// C.(*p).name
// D.p->name
// 参考答案及解析：AC。& 取址运算符，* 指针解引用。
type MyInt1 int
type MyInt2 = int

func Demo1() {
	var i int = 0
	// var i1 MyInt1 = i
	//编译不通过，cannot use i (type int) as type MyInt1 in assignment。
	//将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过
	var i1 MyInt1 = MyInt1(i) //使用强制类型转化 var i1 MyInt1 = MyInt1(i)

	var i2 MyInt2 = i
	//注意类型别名的定义时用了 = ,MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
	fmt.Println(i, i1, i2)
}
