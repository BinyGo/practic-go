package main

import "fmt"

func Demo1() {
	n := 43210
	fmt.Println(n/60*60, " hours and ", n%60*60, " seconds")
	//参考答案及解析：43200 hours and 600 seconds。知识点：运算符优先级。算术运算符 *、/ 和 % 的优先级相同，从左向右结合。

	//修复代码如下：
	fmt.Println(n/(60*60), "hours and", n%(60*60), "seconds")
}

const (
	Century = 100
	Decade  = 010
	Year    = 001
)

func Demo2() {
	fmt.Println(Century + 2*Decade + 2*Year)
	//参考答案及解析：118。知识点：进制数。Go 语言里面，八进制数以 0 开头，十六进制数以 0x 开头，所以 Decade 表示十进制的 8。
}
