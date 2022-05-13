package day059

import "fmt"

type N int

func (n *N) test() {
	fmt.Println(*n)
}

func Demo1() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n) //13
	f1()           //13
	f2()           //13
	//参考答案及解析：13 13 13。知识点：方法值。当目标方法的接收者是指针类型时，那么被复制的就是指针。
}

func Demo2() {
	var m map[int]bool
	_ = m[123]
	a, ok := m[333]
	fmt.Println(a, ok) //false false

	var p *[5]string
	for range p {
		_ = len(p)
	}

	var s []int
	_ = s[:]
	s, s[0] = []int{1, 2}, 9 //runtime error: index out of range [0] with length 0
	fmt.Println(s)
}
