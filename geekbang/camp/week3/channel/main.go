package main

import (
	"fmt"
	"sync"
	"time"
)

func UnBuffered() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- `foo`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)
		println(`Message: ` + <-c)
	}()

	wg.Wait()
}

func Buffered() {
	c := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- `foo`
		c <- `bar`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)
		println(`Message: ` + <-c)
		println(`Message: ` + <-c)
	}()

	wg.Wait()
}

//在这种情况下，Go 只是将语句替换为select以非阻塞模式读取通道
func Demo1() {
	t := time.NewTicker(time.Second * 10)
	select {
	case <-t.C:
		fmt.Println("10 second later...")
	default:
		fmt.Println("default branch")
		//如果代码没有提供default case，Go 会通过阻塞通道
	}
}

func Demo2() {
	ch := make(chan string, 1)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println("ch")
	case <-timeout:
		fmt.Println("timeout")
		ch <- "biny"
	}
	//timeout
}
