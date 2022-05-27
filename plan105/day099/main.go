package main

import "fmt"

// 1.下面代码能编译通过吗？
func Demo1() {
	true := false
	fmt.Println(true) //false
	//参考答案即解析：编译通过。true 是预定义标识符可以用作变量名，但是不建议这么做。
}

//2.下面的代码输出什么？

func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func Demo2() {
	result := watShadowDefer(50)
	fmt.Println(result)
}
