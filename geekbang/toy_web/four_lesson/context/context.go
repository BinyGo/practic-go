package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	WithTimeout()
	WithCancel()
	WithDeadline()
	WithValue()
}

func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	start := time.Now().Unix()
	<-ctx.Done()
	end := time.Now().Unix()
	//输出2,说明ctx.Done()这里阻塞了两秒
	fmt.Println(end - start)
}

func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
		fmt.Println("context wa canceled")
	}()
	// 确保我们的 goroutine进去执行了
	time.Sleep(time.Second)
	cancel()
	// 确保后面那句打印出来了
	time.Sleep(time.Second)
}

func WithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()

	start := time.Now().Unix()
	<-ctx.Done()
	end := time.Now().Unix()
	// 输出2，说明在 ctx.Done()这里阻塞了两秒
	fmt.Println(end - start)
}

func WithValue() {
	parentKey := "parent"
	parent := context.WithValue(context.Background(), parentKey, "this is parent")

	sonKey := "son"
	son := context.WithValue(parent, sonKey, "this is son")

	if parent.Value(sonKey) == nil {
		fmt.Println("parent can not get son's key-value pair")
	}

	if val := son.Value(parentKey); val != nil {
		fmt.Println("son can not get parent key-value pair")
	}
}
