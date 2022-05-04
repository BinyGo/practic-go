package day032

import "fmt"

//1.下面这段代码输出结果正确吗？
type Foo struct {
	bar string
}

func Demo1() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2]) //{A} {B} {C}
	fmt.Println(s2[0], s2[1], s2[2]) //&{C} &{C} &{C}
	// 参考答案及解析：s2 的输出是 &{C} &{C} &{C}，在 第 30 天 的答案解析第二题，我们提到过，for range 使用短变量声明(:=)的形式迭代变量时，
	// 变量 i、value 在每次循环体中都会被重用，而不是重新声明。所以 s2 每次填充的都是临时变量 value 的地址，而在最后一次循环中，value 被赋值为{c}。因此，s2 输出的时候显示出了三个 &{c}。
	for i := range s1 {
		s2[i] = &s1[i]
	}
	fmt.Println(s1[0], s1[1], s1[2]) //{A} {B} {C}
	fmt.Println(s2[0], s2[1], s2[2]) //&{A} &{B} &{C}
}

//2.下面代码里的 counter 的输出值？
func Demo2() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
	// 参考答案及解析：C。for range map 是无序的，如果第一次循环到 A，则输出 3；否则输出 2。
}
