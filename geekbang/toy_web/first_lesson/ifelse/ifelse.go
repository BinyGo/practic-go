package main

import "fmt"

func main() {
	Young(16)
	IfUsingNewVariable(200, 100)
}

func Young(age int) {
	if age < 18 {
		fmt.Println("I am a child!")
	} else {
		// else 分支也可以没有
		fmt.Println("I not a child!")
	}
}

func IfUsingNewVariable(start int, end int) {
	if distance := end - start; distance > 100 {
		fmt.Printf("距离太远,不来了: %d", distance)
	} else {
		fmt.Printf("距离不远,来一趟: %d", distance)
	}

	// 这里不能访问  distance ,作用域仅在 ifelse 里
	//fmt.Printf("距离是： %d\n", distance)
}
