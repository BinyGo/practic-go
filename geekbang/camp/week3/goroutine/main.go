package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// keep yourself busy or do the work yourself
// 当前goroutine需要等另外一个goroutine工作完成,等待它结果返回,才能继续执行(一般需要用waitGroup或者channel来实现)
// 在这种情况下直接在当前goroutine完成这个工作是最简单的,因为委派出去,需要等待另外的goroutine完成,再回调回来

func V1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello,biny")
	})

	//子goroutine,当子goroutine出现err\panic等情况退出,主子goroutine无法感知到
	go func() {
		if err := http.ListenAndServe(":8999", nil); err != nil {
			log.Fatal(err)
			//log.Fatal里会调用os.Exit(1),会无条件终止程序,defers不会被调用到
			//不建议在生产环境下代用log.Fatal,仅在init或者初始化必要配置时考虑使用
		}
	}()

	//空select永远阻塞(pending)在这个地方
	select {}
}

func V2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hello,biny")
	})
	//启动2个端口,2个都会pending,只能一个在子goroutine下执行
	go http.ListenAndServe("127.0.0.1:8888", http.DefaultServeMux)
	//主线程在pending在一个端口
	http.ListenAndServe("0.0.0.0:8999", mux)
}

//V3问题,serveApp退出,整个线程退出,serveDebug退出无法感知到
func V3() {
	go serveDebug()
	serveApp()
}

func serveApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hello,biny")
	})
	http.ListenAndServe("0.0.0.0:8999", mux)
}

func serveDebug() {
	http.ListenAndServe("127.0.0.1:8888", http.DefaultServeMux)
}

//V4 还是有各自为政的问题
func V4() {
	go serveAppV4()
	go serveDebugV4()
	select {}
}

func serveAppV4() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hello,biny")
	})
	if err := http.ListenAndServe("0.0.0.0:8999", mux); err != nil {
		log.Fatal(err)
	}
}

func serveDebugV4() {
	if err := http.ListenAndServe("127.0.0.1:8888", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

//让调用者来决定要不要并发,让使用权交给使用者
func V5() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	//启动一个serveApp
	go func() {
		done <- serveV5(
			"127.0.0.1:8888",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(w, "hello,biny")
			},
			stop,
		)
	}()
	//启动一个serveDebug
	go func() {
		done <- serveV5(
			"0.0.0.0:8999",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(w, "hello,debug")
			},
			stop,
		)
	}()

	var stopped bool
	//通过for循环监听done
	for i := 0; i < cap(done); i++ {
		//当有一个子线程done退出时,打印日志,调用close(stop)
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}

func serveV5(addr string, handler http.HandlerFunc, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		//监听主线程stop chan,当close(stop),执行退出,让外面可以控制退出
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

// Leave concurrency to the caller
func TrackerExec() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")
	time.Sleep(time.Second * 3)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	tr.Shutdown(ctx)
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(time.Second * 1)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}
