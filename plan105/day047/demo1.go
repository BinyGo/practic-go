package day047

import "fmt"

func Demo1() {
	data := []int{1, 2, 3}
	i := 0
	//++i //expected statement, found '++'
	//fmt.Println(data[i++]) //expected ']', found '++'
	// 对于自增、自减，需要注意：
	// 自增、自减不在是运算符，只能作为独立语句，而不是表达式；
	// 不像其他语言，Go 语言中不支持 ++i 和 –i 操作；
	// 表达式通常是求值代码，可作为右值或参数使用。而语句表示完成一个任务，比如 if、for 语句等。表达式可作为语句使用，但语句不能当做表达式。
	i++
	fmt.Println(data[i]) //2
}

//2.下面代码最后一行输出什么？请说明原因。
func Demo2() {
	x := 1
	fmt.Println(x) //1
	{
		fmt.Println(x) //1
		i, x := 2, 2
		fmt.Println(i, x) //2,2
		x = 2
	}
	fmt.Println(x) //1
	// 输出1。知识点：变量隐藏。使用变量简短声明符号 := 时，
	// 如果符号左边有多个变量，只需要保证至少有一个变量是新声明的，并对已定义的变量尽进行赋值操作。
	// 但如果出现作用域之后，就会导致变量隐藏的问题，就像这个例子一样。
}
