package day055

import "fmt"

//1.下面的代码有什么问题？
type T struct {
	n int
}

func getT() T {
	return T{}
}

func Demo1() {
	//getT().n = 1 //cannot assign to getT().n (value of type int)
	//直接返回的 T{} 无法寻址，不可直接赋值。
	t := getT()
	t.n = 1
	p := &t.n
	*p = 2
	fmt.Println(t.n) //2
}
