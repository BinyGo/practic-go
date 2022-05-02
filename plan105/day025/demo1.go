package day025

import "fmt"

/*
1.下面这段代码输出什么？为什么？
func (i int) PrintInt() {
	fmt.Println(i)
}

func Demo1() {
	var i int = 1
	i.PrintInt()
}
*/
// A. 1
// B. compilation error
// 参考答案及解析：B。基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。
// 解决的办法可以定义一种新的类型：
type MyInt int

func (i MyInt) PrintInt() {
	fmt.Println(i)
}

func Demo2() {
	var i MyInt = 1
	i.PrintInt()
}

//2.下面这段代码输出什么？为什么？

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "Speak" {
		talk = "Speak"
	} else {
		talk = "hi"
	}
	return
}

func Demo3() {
	//var peo People = Student{} //cannot use (Student literal) (value of type Student) as People value in variable declaration: missing method Speak (Speak has pointer receiver)
	var peo People = &Student{}
	think := "Speak"
	fmt.Println(peo.Speak(think))
}

// var peo People = Student{}
// 编译错误 Student does not implement People (Speak method has pointer receiver)
// 值类型 Student 没有实现接口的 Speak() 方法，而是指针类型 *Student 实现该方法。
