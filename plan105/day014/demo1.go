package day014

import "fmt"

//1.下面代码输出什么？
func Demo1() {
	str := "hello"
	//str[0] = "x" //cannot assign to str[0] (value of type byte)
	fmt.Println(str)
	// 	A. hello
	// B. xello
	// C. compilation error
	// 参考代码及解析：C。知识点：常量，Go 语言中的字符串是只读的。
}

func incr(p *int) int {
	*p++
	return *p
}

//2.下面代码输出什么？
func Demo2() {
	p := 1
	incr(&p)
	// 	fmt.Println(p)
	// 	A. 1
	// 	B. 2
	//  C. 3
	// 	参考答案及解析：B。知识点：指针，incr() 函数里的 p 是 *int 类型的指针，指向的是Demo2函数变量 p 的地址。第 2 行代码是将该地址的值执行一个自增操作，incr() 返回自增后的结果。
}

//3.对 add() 函数调用正确的是（）
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum

}

func DemoAdd() {
	add(1, 2)
	add(1, 3, 7)
	//add([]int{1, 2}) //cannot use ([]int literal) (value of type []int) as int value in argument to add
	add([]int{1, 3, 7}...)
	// A. add(1, 2)
	// B. add(1, 3, 7)
	// C. add([]int{1, 2})
	// D. add([]int{1, 3, 7}…)
	// 参考答案及解析：ABD。知识点：可变函数。
}
