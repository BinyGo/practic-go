package day046

import (
	"fmt"
)

const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func Demo1() {
	fmt.Printf("%T %v \n", y, y) //uint16 120
	fmt.Printf("%T %v \n", z, z) //string abc
	//常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
}

func Demo2() {
	//var x string = nil //cannot use nil (untyped nil value) as string value in variable declaration
	var x string
	if x == "" {
		x = "default"
	}
	//不能将 nil 分配给 string 类型的变量。
}
