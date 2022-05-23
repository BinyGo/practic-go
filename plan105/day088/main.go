package main

import (
	"fmt"
	"os"
)

// 1.下面这段代码能通过编译吗？请简要说明。
func Demo1() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}

// 参考答案及解析：能通过编译。
// 上面的代码可以理解成：
func Demo1_2() {
	m := make(map[string]int)
	v := m["foo"] //下标不存在返回类型零值,要判断是否存在下标使用comm,ok
	v++
	m["foo"] = v
	fmt.Println(m["foo"])
}

func Foo() error {
	var err *os.PathError = nil
	// …
	return err
}

func Demo2() {
	err := Foo()
	fmt.Println(err)        //<nil>
	fmt.Println(err == nil) //false
	// 参考答案及解析：nil false。知识点：接口值与 nil 值。只有在值和动态类型都为 nil 的情况下，接口值才为 nil。
	// Foo() 函数返回的 err 变量，值为 nil、动态类型为 *os.PathError，与 nil（值为 nil，动态类型为 nil）显然是不相等。我们可以打印下变量 err 的详情：
	fmt.Printf("%#v\n", err) // (*fs.PathError)(nil)
}
