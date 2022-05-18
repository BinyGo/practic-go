package day073

//1.下面代码输出什么？
func test() []func() {
	var funcs []func()
	for i := 0; i < 2; i++ {
		funcs = append(funcs, func() {
			println(&i, i)
		})
	}
	return funcs
}

func Demo1() {
	funs := test()
	for _, f := range funs {
		f()
	}
	// 0xc0000be070 2
	// 0xc0000be070 2
	//知识点：闭包延迟求值。for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。
}

//2.下面的代码能编译通过吗？可以的话输出什么，请说明？
var f = func(i int) {
	print("x")
}

func Demo2() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
	// 参考答案及解析：10x。这道题一眼看上去会输出 109876543210，其实这是错误的答案，这里不是递归。
	// 假设 main() 函数里为 f2()，外面的为 f1()，当声明 f2() 时，调用的是已经完成声明的 f1()。
}

//下面这段代码会更容易理解一点：

var x = 23

func Demo3() {
	x := 2*x - 4
	println(x) //42
}
