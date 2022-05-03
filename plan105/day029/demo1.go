package day029

import (
	"fmt"
	"time"
)

//1.下面这段代码能否正常结束？
func Demo1() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
	// 参考答案及解析：不会出现死循环，能正常结束。
	// 循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。
}

//2.下面这段代码输出什么？为什么？

func Demo2() {
	fmt.Println("Demo2")
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 1)
	// 2 3
	// 2 3
	// 2 3
}

func Demo3() {
	fmt.Println("Demo3")
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
		time.Sleep(time.Second * 1)
	}
	time.Sleep(time.Second * 1)
	// 0 1
	// 1 2
	// 2 3
}

// for range 使用短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明。

// 各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值，而不是各个goroutine启动时的i, v值。可以理解为闭包引用，使用的是上下文环境的值。

// 两种可行的 fix 方法:

func Demo4() {
	fmt.Println("Demo4")
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 1)
	// 2 3
	// 1 2
	// 0 1
	// 乱序
}

func Demo5() {
	fmt.Println("Demo5")
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		i := i
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 1)
	// 2 3
	// 1 2
	// 0 1
	// 乱序
}
