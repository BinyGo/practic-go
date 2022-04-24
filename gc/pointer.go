//Go GC如何检测内存对象中是否包含指针
//https://mp.weixin.qq.com/s/uVhRqQh0gwzUeP93UXJ5wA
package gc

import (
	"fmt"
	"runtime"
	"time"
)

//GC只关心指针，只要被扫描到的内存对象中有指针，它就会“顺藤摸瓜”，
//把该内存对象所在的“关系网”摸个门儿清，而那些被孤立在这张“网”之外的内存对象就是要被“清扫”的对象。

func GcDemo1() {
	//demo1中声明了一个包含10亿个*int的切片变量a
	a := make([]*int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		//调用runtime.GC函数手工触发GC过程，并度量每次GC的执行时间
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
		// GC took 235.542882ms
		// GC took 118.721986ms...
	}

	runtime.KeepAlive(a)
	//程序中调用runtime.KeepAlive函数用于保证在该函数调用点之前切片a不会被GC释放掉。
}

func GcDome2() {
	//将切片的元素类型由*int改为了int
	a := make([]int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
		// GC took 982.436µs
		// GC took 381.271µs...
	}

	runtime.KeepAlive(a)
}

//实验环境中demo2中每轮GC的性能是demo1的300多倍！
//两个demo源码唯一的不同就是切片中的元素类型，demo1中的切片元素类型为int型指针。
//GC每次触发后都会全量扫描切片中存储的这10亿个指针，这就是demo1 GC函数执行时间很长的原因。
//而demo2中的切片元素类型为int，从demo2的运行结果来看，GC根本没有搭理demo2中的a，这也是demo2 GC函数执行时间较短的原因
//通过以上GC行为差异，我们知道GC可以通过切片a的类型知晓其元素是否包含指针，进而决定是否对其进行进一步扫描。
