package day074

import (
	"fmt"
	"runtime"
)

func Demo1() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
	} //this loop will spin, using 100% CPU (SA5002)
	//参考答案及解析：for {} 独占 CPU 资源导致其他 Goroutine 饿死。
}

//可以通过阻塞的方式避免 CPU 占用，修复代码：
func Demo2() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		//os.Exit(0)
	}()
	select {}
}
