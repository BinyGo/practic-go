package main

import (
	"fmt"
	"time"
)

func main() {
	i := 13
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	a()

	fmt.Println(ReturnClosure("Biny")())

	Delay()
	time.Sleep(time.Second)
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "hello, " + name
	}
}

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		i := i
		fns = append(fns, func() {
			fmt.Printf("hello,this is : %d \n", i)
		})
	}
	for _, fn := range fns {
		fn()
	}
}
