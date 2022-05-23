package main

import (
	"fmt"
)

// 1.下面代码能编译通过吗？请简要说明。

func Demo1() {
	v := []int{1, 2, 3}
	for i, n := 0, len(v); i < n; i++ {
		v = append(v, i)
	}
	fmt.Println(v)
	//参考答案及解析：能编译通过，输出 [1 2 3 0 1 2]。for 循环开始的时候，终止条件只会计算一次。
}

type P *int
type Q *int

func Demo2() {
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++
	fmt.Println(*p, *q)
	// A.8 8
	// B.8 9
	// C.9 9
	// 参考答案及解析：C。指针变量指向相同的地址。
}
