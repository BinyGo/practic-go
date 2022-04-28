package day013

import "fmt"

// 1.定义一个包内全局字符串变量，下面语法正确的是（）
// A. var str string
// B. str := ""
// C. str = ""
// D. var str = ""
// 参考答案及解析：AD。B 只支持局部变量声明；C 是赋值，str 必须在这之前已经声明；

func hello(i int) {
	fmt.Println("hello:", i)
}

func Demo1() {
	i := 5
	defer hello(i) //hello:5 hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.
	i = i + 10
	fmt.Println("demo1:", i) //demo1:15
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("ShowA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teatcher showB")
}

func Demo2() {
	t := Teacher{}
	t.ShowA()
	//showA
	//showB
	//知识点：结构体嵌套。这道题可以结合第 12 天的第三题一起看，Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法，需要注意的是第 5 行代码调用的是 People 自己的 ShowB 方法。
}
