package day015

import (
	"fmt"
	"strconv"
)

//1.下面这段代码输出什么？
func Demo1() {
	var s1 []int
	var s2 = []int{}

	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 not nil")
	}
	//s1 is nil

	if s2 == nil {
		fmt.Println("s2 is nil")
	} else {
		fmt.Println("s2 not nil")
	}
	//s2 not nil

	//知识点：nil 切片和空切片。nil 切片和 nil 相等，一般用来表示一个不存在的切片；空切片和 nil 不相等，表示一个空的集合。
}

//2.下面这段代码输出什么？
func Demo2() {
	i := 65
	//str := string(i) //从int到字符串的转换得到的是一个符文的字符串，而不是一个数字的字符串,运行报错
	str := strconv.Itoa(i)
	fmt.Println(fmt.Sprint(i)) //65
	fmt.Println(str)           //65

}
