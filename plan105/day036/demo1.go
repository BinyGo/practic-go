package day036

import (
	"fmt"
	"runtime"
)

//1.下面代码会触发异常吗？请说明。
func Demo1() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
	//参考答案及解析：select 会随机选择一个可用通道做收发操作，所以可能触发异常，也可能不会。
	//biny:case是有序的,但chan不是有序的
}
