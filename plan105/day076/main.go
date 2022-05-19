package day076

import "fmt"

type S1 struct{}

func (s1 S1) f() {
	fmt.Println("S1.f()")
}

func (s1 S1) g() {
	fmt.Println("s1.g()")
}

type S2 struct {
	S1
}

func (s2 S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I) {
	fmt.Printf("%T\n", i)
	if s1, ok := i.(S1); ok {
		s1.f()
		s1.g()
	}
	if s2, ok := i.(S2); ok {
		s2.f()
		s2.g()
	}
}
func Demo1() {
	printType(S1{})
	printType(S2{})

	// 参考答案及解析：
	// day076.S1
	// S1.f()
	// s1.g()
	// day076.S2
	// S2.f()
	// s1.g()
	//知识点：类型断言，结构体嵌套。结构体 S2 嵌套了结构体 S1，S2 自己没有实现 g() ，调用的是 S1 的 g()。
}
