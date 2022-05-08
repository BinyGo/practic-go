package day045

import "fmt"

func Demo1() {
	one := 0
	one, two := 1, 2
	one, two = two, one
	fmt.Println(one, two)
}

func test(x byte) {
	fmt.Println(x)
}

func Demo2() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	test(a) //17
	test(b) //17
	test(c) //34
	//与 rune 是 int32 的别名一样，byte 是 uint8 的别名，别名类型无序转换，可直接转换。
}
