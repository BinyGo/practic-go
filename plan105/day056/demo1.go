package day056

//1.下面的代码有什么问题？

import "fmt"

func Demo1() {
	s := make([]int, 3, 9)
	fmt.Println(len(s)) //3
	s2 := s[4:8]
	fmt.Println(len(s2)) //4
	// 从一个基础切片派生出的子切片的长度可能大于基础切片的长度。
	// 假设基础切片是 baseSlice，使用操作符 [low,high]，
	// 有如下规则：0 <= low <= high <= cap(baseSlice)，
	// 只要上述满足这个关系，下标 low 和 high 都可以大于 len(baseSlice)。
}

//2.下面代码输出什么？
type N int

func (n N) test() {
	fmt.Println(n)
}
func Demo2() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test
	fmt.Println(n) //13
	f1()           //11
	f2()           //12
	// 参考答案及解析：13 11 12。知识点：方法值。
	// 当指针值赋值给变量或者作为函数参数传递时，
	// 会立即计算并复制该方法执行所需的接收者对象，
	// 与其绑定，以便在稍后执行时，能隐式第传入接收者参数。
}
