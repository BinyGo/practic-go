package day030

import "fmt"

//1.下面这段代码输出什么？

func f(n int) (r int) {
	defer func() {
		fmt.Println(r) //4
		fmt.Println(n) //3
		r += n
		recover()
	}()

	var f func()
	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func Demo1() {
	fmt.Println(f(3))
}

// 参考答案及解析：7。根据 5 年 Gopher 都不知道的 defer 细节，你别再掉进坑里！
// 提到的“三步拆解法”，第一步执行r = n +1，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return。
// 此题引自知识星球《Go项目实战》。

//2.下面这段代码输出什么？
func Demo2() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	// 参考答案及解析：
	// r =  [1 2 3 4 5]
	// a =  [1 12 13 4 5]
	// range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a
}

//如果想要 r 和 a 一样输出，for range &a：
func Demo3() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	//副本是一个指向原数组 a 的指针，因此后续所有循环中均是 &a 指向的原数组亲自参与的，因此 v 能从 &a 指向的原数组中取出 a 修改后的值。
}
