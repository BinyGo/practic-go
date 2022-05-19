package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"sync"
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
	time.Sleep(1 * time.Second)
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

// 外部for Client 3次,确是3条 request from client 7,time: 1652953141026362
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

//内部100个并发只有一条 request from client 7,time: 1652953141026362
func Client2() {
	client, err := rpc.DialHTTP("tcp", ":8999")
	if err != nil {
		panic(err)
	}

	singleFly := new(singleflight.Group)
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
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
