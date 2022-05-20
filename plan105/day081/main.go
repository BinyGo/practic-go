package day081

import "fmt"

func Demo1() {
	var a []int = nil
	a, a[0] = []int{1, 2}, 9 //runtime error: index out of range [0] with length 0
	fmt.Println(a)

	// 参考答案即解析：运行时错误。知识点：多重赋值。
	// 多重赋值分为两个步骤，有先后顺序：
	// 计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；这时左边的a[0]不存在,对其赋值会panic
	// 赋值；
}

//2.下面代码中的指针 p 为野指针，因为返回的栈内存在函数结束时会被释放？
type TimeMatcher struct {
	base int
}

func NewTimeMaster(base int) *TimeMatcher {
	return &TimeMatcher{base: base}
}

func Demo2() {
	p := NewTimeMaster(3)
	fmt.Println(p) //&{3}

	// A. false
	// B. true
	// 参考答案及解析：A。Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会被释放（内存逃逸），因此在 Go 语言中返回函数参数或临时变量是安全的。
}
