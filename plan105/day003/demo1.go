package day003

import "fmt"

//1.下面两段代码输出什么。
func Demo1() {

	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	//[0 0 0 0 0 1 2 3]

	c := make([]int, 0)
	c = append(c, 1, 2, 3, 4)
	fmt.Println(c)
	//[1 2 3 4]

}

// 2.new() 与 make() 的区别
// 参考答案：

// new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

// new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

// make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
