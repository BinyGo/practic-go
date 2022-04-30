package day021

import "fmt"

// 1.下面的两个切片声明中有什么区别？哪个更可取？
// A. var a []int
// B. a := []int{}
// 参考答案及解析：A 声明的是 nil 切片；B 声明的是长度和容量都为 0 的空切片。第一种切片声明不会分配内存，优先选择。

//2. A、B、C、D 哪些选项有语法错误？
type S struct {
}

func f(x interface{}) {
}

// func g(x *interface{}) {
// }

func Demo1() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B
	f(p) //C
	//g(p) //D
	// 	参考答案及解析：BD。函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也用 interface{}，而不是使用 *interface{}。
	// 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
}

//3.下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？
type DS struct {
	m string
}

func DS3() *DS {
	return &DS{"ds3"}
	//函数返回参数是指针类型，所以可以用 & 取结构体的指针；
}

func Demo3() {
	p := DS3()
	p2 := *DS3()
	fmt.Println(p.m)
	fmt.Println(p2.m)
	//如果填 *DS3()，则 p 是 DS 类型；如果填 DS3()，则 p 是 *DS 类型，不过都可以使用 p.m 取得结构体的成员。
}
