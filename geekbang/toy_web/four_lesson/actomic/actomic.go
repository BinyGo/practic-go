package main

import "sync/atomic"

var value int32 = 0

func main() {
	//要传入的value的指针
	//把value+10
	atomic.AddInt32(&value, 10)
	nv := atomic.LoadInt32(&value)
	println(nv) //10
	// 如果之前的值是10，那么就设置为新的值 20
	swapped := atomic.CompareAndSwapInt32(&value, 10, 20)
	println(swapped) //true

	// 如果之前的值是19，那么就设置为新的值 50
	// 显然现在 value 是 20
	swapped = atomic.CompareAndSwapInt32(&value, 19, 50)
	println(swapped) //false

	old := atomic.SwapInt32(&value, 40)
	println(old)   // 20，即原本的值
	println(value) // 40 输出新的值，也就是交换后的值
}
