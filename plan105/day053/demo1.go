package day053

import "fmt"

// 1.关于 channel 下面描述正确的是？
// A. 向已关闭的通道发送数据会引发 panic；

// B. 从已关闭的缓冲通道接收数据，返回已缓冲数据或者零值；

// 参考答案及解析：AB。

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}

func getT() T {
	return T{}
}

func Demo1() {
	//getT().Set(1) //cannot call pointer method Set on T
	// 有两处问题：
	// 1.直接返回的 T{} 不可寻址；
	// 2.不可寻址的结构体不能调用带结构体指针接收者的方法；
	t := getT()
	t.Set(2)
	fmt.Println(t.n)
}
