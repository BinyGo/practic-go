package day039

import (
	"fmt"
	"time"
)

// 1.关于无缓冲和有冲突的channel，下面说法正确的是？
// A. 无缓冲的channel是默认的缓冲为1的channel；
// B. 无缓冲的channel和有缓冲的channel都是同步的；
// C. 无缓冲的channel和有缓冲的channel都是非同步的；
// D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的；
// 参考答案及解析：D。

//2.下面代码是否能编译通过？如果通过，输出什么？
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func Demo1() {
	var x *int = nil
	Foo(x) //non-empty interface
	// 参考答案及解析：non-empty interface 考点：interface 的内部结构，
	// 我们知道接口除了有静态类型，还有动态类型和动态值，当且仅当动态值和动态类型都为 nil 时，
	// 接口类型值才为 nil。这里的 x 的动态类型是 *int，所以 x 不为 nil。
}

//3.下面代码输出什么？
func Demo2() {
	ch := make(chan int, 100)
	//A
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	//B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
	// ok
	// close
	// panic: send on closed channel
	// 参考答案及解析：程序抛异常。先定义下，第一个协程为 A 协程，第二个协程为 B 协程；
	// 当 A 协程还没起时，主协程已经将 channel 关闭了，当 A 协程往关闭的 channel 发送数据时会 panic，panic: send on closed channel。
}
