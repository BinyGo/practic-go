package day063

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func Demo1() {
	f := F(5)
	defer func() {
		fmt.Println(f()) //8
	}()
	defer fmt.Println(f()) //6
	i := f()
	fmt.Println(i) //7
	//参考答案及解析：768。知识点：匿名函数、defer()。defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出。
}
