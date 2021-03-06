package main

import (
	"fmt"
)

func main() {
	name := "Tom"
	age := 17
	// 这个 API 是返回字符串的，所以大多数时候我们都是用这个
	str := fmt.Sprintf("hello, I am %s, I am %d years old", name, age)
	println(str)

	// 这个是直接输出，一般简单程序 DEBUG 会用它输出到一些信息到控制台
	fmt.Printf("hello, I am %s, I am %d years old \n", name, age)

	replaceHolder()
}

type user struct {
	Name string
	Age  int
}

func replaceHolder() {
	u := &user{
		Name: "Biny",
		Age:  18,
	}
	fmt.Printf("v => %v \n", u)   //v => &{Biny 18}
	fmt.Printf("+v => %+v \n", u) //+v => &{Name:Biny Age:18}
	fmt.Printf("#v => %#v \n", u) //#v => &main.user{Name:"Biny", Age:18}
	fmt.Printf("T => %T \n", u)   //T => *main.user
}
