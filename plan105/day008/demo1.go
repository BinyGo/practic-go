package day008

import "fmt"

/*
1.关于init函数，下面说法正确的是()
A. 一个包中，可以包含多个 init 函数；
B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
C. main 包中，不能有 init 函数；
D. init 函数可以被其他函数调用；
1.参考答案及解析：AB。关于 init() 函数有几个需要注意的地方：

init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;
一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；
同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的;
init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；
引入包，不可出现死循坏。即 A import B,B import A，这种情况编译失败；
*/

/*
下面这段代码能否编译通过？如果可以，输出什么？
func GetValue() int {
	return 1
}
编译失败。考点：类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择。
*/
func GetValue() interface{} {
	return 1
}

func Demo1() {
	i := GetValue()
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}
