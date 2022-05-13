package day060

import (
	"fmt"
	"sync"
	"time"
)

type T struct{}

func (*T) foo() {}

func (T) bar() {}

type S struct {
	*T
}

func Demo1() {
	s := S{}
	_ = s.foo
	s.foo()
	//_ = s.bar
	//参考答案及解析：第 17 行，因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic。
}

type data struct {
	sync.Mutex
}

func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func (d *data) test2(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func Demo2() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
	//参考答案及解析：锁失效。将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。
}

func Demo3() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data

	go func() {
		defer wg.Done()
		d.test2("read")
	}()

	go func() {
		defer wg.Done()
		d.test2("write")
	}()
	wg.Wait()
}
