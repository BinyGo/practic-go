package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"time"

	"golang.org/x/sync/singleflight"
)

type (
	Arg struct {
		Caller int
	}
	Data struct{}
)

var singleFly = new(singleflight.Group)

func (d *Data) GetData(arg *Arg, replay *string) error {
	key := "biny"
	result, _, shared := singleFly.Do(key, func() (interface{}, error) {
		err := d.getData(arg, replay)
		return replay, err
	})
	fmt.Printf("get result: %#v, shared: %v\n", result, shared)
	if *replay == "" {
		d.getRedis(arg, replay)
	}
	return nil

}

func (d *Data) getData(arg *Arg, replay *string) error {
	fmt.Printf("request from client %d,time: %d\n", arg.Caller, time.Now().UnixMicro())
	*replay = fmt.Sprintf("source data from rpcServer,time:%d", time.Now().UnixMicro())
	time.Sleep(time.Second)
	return nil
}

func (d *Data) getRedis(arg *Arg, replay *string) error {
	time.Sleep(time.Millisecond * 100)
	//延时一下,再从redis获取或者降级处理
	*replay = "redis获取或者降级处理"
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

func Client() {
	client, err := rpc.DialHTTP("tcp", ":8999")
	if err != nil {
		panic(err)
	}

	var result string
	err = client.Call("Data.GetData", Arg{Caller: 1}, &result)
	fmt.Printf("get result: %s\n", result)

	if err != nil {
		return
	}

}
