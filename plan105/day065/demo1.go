package day065

import "fmt"

// 1.flag 是 bool 型变量，下面 if 表达式符合编码规范的是？
// A. if flag == 1
// B. if flag
// C. if flag == false
// D. if !flag
// 参考答案及解析：BCD。

func Demo1() {
	defer func() {
		fmt.Println("a:", recover())
	}()
	defer func() {
		defer func() {
			fmt.Println("b:", recover())
		}()
		panic(1)
	}()
	defer recover() //无效
	panic(2)
	// b: 1
	// a: 2
	// 相关知识点请看 第64天 题目解析。
}
