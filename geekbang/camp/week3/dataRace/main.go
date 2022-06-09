package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
var Counter int = 0

// go build -race
// go run -race main.go
// WARNING: DATA RACE
func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}
	Wait.Wait()
	fmt.Printf("Final Counter: %d \n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter //race:Read at...main.go:24 +0x39 产生一个读race
		value++
		Counter = value //write:Previous write at...main.go:26 +0x39 产生一个写race
		//修改为Counter=Counter+1或Counter++,还是会出现 读写race
		//一个i++操作并不是一个原子操作,汇编后一般生成3条指令
		//需用互斥锁或者atomic方式
	}
	Wait.Done()
}
