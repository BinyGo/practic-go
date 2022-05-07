package day040

import "fmt"

// 1.关于select机制，下面说法正确的是?
// A. select机制用来处理异步IO问题；
// B. select机制最大的一条限制就是每个case语句里必须是一个IO操作；
// C. golang在语言级别支持select关键字；
// D. select关键字的用法与switch语句非常类似，后面要带判断条件；
// 参考答案及解析：ABC。

//2.下面的代码有什么问题？
func Stop(stop <-chan bool) {
	//close(stop) //invalid argument: stop (variable of type <-chan bool) must not be a receive-only channel
	//参考答案及解析：有方向的 channel 不可以被关闭。
}

type Param map[string]interface{}

type Show struct {
	*Param
}

//3.下面这段代码存在什么问题？
func Demo1() {
	//s := new(Show)
	//s.Param["day"] = 2 //invalid operation: cannot index s.Param (variable of type *Param)
	//参考答案及解析：存在两个问题：1.map 需要初始化才能使用；2.指针不支持索引。
}

func Demo2() {
	s := new(Show)
	p := make(Param)
	p["day"] = 2
	s.Param = &p
	//fmt.Println(s.Param["day"]) //invalid operation: cannot index s.Param (variable of type *Param)
	tmp := *s.Param
	fmt.Println(tmp["day"])
}
