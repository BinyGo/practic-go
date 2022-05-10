package main

import "fmt"

//1.下面这段代码输出什么？

type T struct {
	ls []int
}

func foo(t T) {
	t.ls[0] = 100
}

func Demo1() {
	var t = T{
		ls: []int{1, 2, 3},
	}
	foo(t)
	fmt.Println(t.ls[0])
	// 	A. 1
	// B. 100
	// C. compilation error
	// 参考答案及解析：B。调用 foo() 函数时虽然是传值，但 foo() 函数中，字段 ls 依旧可以看成是指向底层数组的指针。
}

//2.下面代码输出什么？

func Demo2() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1)) //false
	fmt.Println(isMatch(2)) //true
	// 参考答案及解析：false true。Go 语言的 switch 语句虽然没有”break”，但如果 case 完成程序会默认 break，
	// 可以在 case 语句后面加上关键字 fallthrough，这样就会接着走下一个 case 语句（不用匹配后续条件表达式）。
	// 或者，利用 case 可以匹配多个值的特性。
}

func Demo3() {
	isMath := func(i int) bool {
		switch i {
		case 1:
			fallthrough
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMath(1)) //true
	fmt.Println(isMath(2)) //true
	match := func(i int) bool {
		switch i {
		case 1, 2:
			return true
		}
		return false
	}
	fmt.Println(match(1)) //true
	fmt.Println(match(2)) //true
}
