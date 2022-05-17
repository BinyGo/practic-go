package day072

import (
	"fmt"
)

//1.下面的代码输出什么，请说明。

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Println(elem)
	return s
}

func Demo1() {
	s := NewSlice()
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
	//参考答案及解析：312。对比昨天的第二题，本题的 s.Add(1).Add(2) 作为一个整体包在一个匿名函数中，会延迟执行。
}

//2.下面的代码输出什么，请说明？

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	fmt.Println("--String()--")
	return fmt.Sprintf("%#v", o.Quantity)
}

func Demo2() {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange.String()) //--String()-- 5
	fmt.Println(orange)          //{5}
	//参考答案及解析：这道题容易忽视的点是，String() 是指针方法，而不是值方法，所以使用 Println() 输出时不会调用到 String() 方法。

}

func Demo3() {
	fmt.Println("--Demo3--")
	orange := &Orange{}
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange.String()) //--String()-- 5
	fmt.Println(orange)          //--String()-- 5
}

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

func Demo4() {
	var t T
	//f1 := T.Set                            // 报错:cannot call pointer method Set on T
	f1 := (*T).Set                           // f1的类型，也是*T类型Set方法的类型：func (t *T, int)int
	f2 := T.Get                              // f2的类型，也是T类型Get方法的类型：func(t T)int
	fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*day072.T, int) int
	fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(day072.T) int
	fmt.Printf("the type of t is %T\n", t)   //the type of t is day072.T

	f1(&t, 3)
	fmt.Println(f2(t)) // 3
}
