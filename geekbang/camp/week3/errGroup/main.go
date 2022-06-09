package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

//errgroup 内部用了WaitGroup来管理goroutine,用sync.Once来做error的处理
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	var a, b, c []int

	//并行3个g.Go
	//调用首页商品信息
	g.Go(func() error {
		fmt.Println("Goods")
		a = nil
		return errors.New("Goods")
	})

	//调用首页AI广告信息
	g.Go(func() error {
		time.Sleep(time.Second * 10)
		fmt.Println("AI")
		b = nil
		return nil
	})

	//调用首页头部信息
	g.Go(func() error {
		fmt.Println("head")
		c = nil
		return nil
	})
	//其他位置,或者日志,trace等等...

	err := g.Wait()
	//wait之后组装a,b,c数据
	fmt.Println(err)
	fmt.Println(ctx.Err())
	fmt.Println(a, b, c)
}
