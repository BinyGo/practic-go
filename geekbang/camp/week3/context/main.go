package main

import (
	"context"
	"fmt"
	"time"
)

func Demo1() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return //returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

const shortDuration = time.Millisecond * 1

func Demo2() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(time.Second * 1):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) //context deadline exceeded
	}
}
