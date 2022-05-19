package main

import (
	"context"
	"fmt"
	"net/http"
	"net/rpc"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

type (
	Arg struct {
		Caller int
	}
	Data struct{}
)

func (d *Data) GetData(arg *Arg, replay *string) error {
	fmt.Printf("request from client %d,time: %d\n", arg.Caller, time.Now().UnixMicro())
	//time.Sleep(1 * time.Second)
	*replay = fmt.Sprintf("source data from rpcServer,time:%d", time.Now().UnixMicro())
	return nil
}

func main() {
	d := new(Data)
	rpc.Register(d)
	rpc.HandleHTTP()
	fmt.Println("start rpc server")
	if err := http.ListenAndServe(":8999", nil); err != nil {
		panic(err)
	}
}

// func Client() {
// 	client, err := rpc.DialHTTP("tcp", ":8999")
// 	if err != nil {
// 		panic(err)
// 	}

// 	singleFly := new(singleflight.Group)
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)

// 	for i := 0; i < 2; i++ {
// 		fn := func() (interface{}, error) {
// 			var replay string
// 			err = client.Call("Data.GetData", Arg{Caller: i}, &replay)

// 			return replay, err
// 		}

// 		go func(i int) {
// 			result, _, shared := singleFly.Do("foo", fn)
// 			fmt.Printf("caller: %d, get result: %s, shared: %v\n", i, result, shared)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

func Client() {
	client, err := rpc.DialHTTP("tcp", ":8999")
	if err != nil {
		panic(err)
	}

	singleFly := new(singleflight.Group)
	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < 2; i++ {
		fn := func() (interface{}, error) {
			var replay string
			err = client.Call("Data.GetData", Arg{Caller: i}, &replay)

			return replay, err
		}

		go func(i int) {
			result, _, shared := singleFly.Do("foo", fn)
			fmt.Printf("caller: %d, get result: %s, shared: %v\n", i, result, shared)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q %d", query, time.Now().UnixMicro())), nil
}

func Demo1() {
	var g singleflight.Group
	const n = 5
	waited := int32(n)
	done := make(chan struct{})
	key := "biny"
	for i := 0; i < n; i++ {
		go func(j int) {
			v, _, shared := g.Do(key, func() (interface{}, error) {
				ret, err := find(context.Background(), key)
				return ret, err
			})
			if atomic.AddInt32(&waited, -1) == 0 {
				close(done)
			}
			fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
		}(n)
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		fmt.Println("Do hangs")
	}
}

func Demo2() {
	singleFly := new(singleflight.Group)
	wg := sync.WaitGroup{}
	wg.Add(30)
	key := "biny"

	fn := func() (interface{}, error) {
		ret, err := find(context.Background(), key)
		return ret, err
	}

	for i := 0; i < 30; i++ {
		go func(i int) {
			result, _, shared := singleFly.Do(key, fn)
			fmt.Printf("caller: %d, get result: %s, shared: %v\n", i, result, shared)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
