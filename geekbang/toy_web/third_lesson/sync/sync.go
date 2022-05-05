package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup()
	pool()
	PrintOnce()
	PrintOnce()
}

func waitGroup() {
	res := 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(val int) {
			res += val
			wg.Done()
		}(i)
	}
	// 把这个注释掉你会发现，不会等待上面结果,直接执行后续代码
	wg.Wait()
	fmt.Println("waitGroup res:", res) //waitGroup res: 45
}

type user struct {
	Name  string
	Email string
}

func (u *user) Reset(name string, email string) {
	u.Email = email
	u.Name = name
}

func pool() {
	pool := sync.Pool{
		New: func() interface{} {
			return &user{}
		},
	}
	// Get 返回的是 interface{}，所以需要类型断言
	u := pool.Get().(*user)
	// defer 还回去
	defer pool.Put(u)

	// 紧接着重置 u 这个对象
	u.Reset("biny", "58040087@qq.com")

	// 下边就是使用 u 来完成你的业务逻辑
}

var once sync.Once

// 这个方法，不管调用几次，只会输出一次
func PrintOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}

var mutex sync.Mutex
var rwMutex sync.RWMutex

func Mutex() {
	mutex.Lock()
	defer mutex.Unlock()
	//要执行的代码,防并发冲突
}
func RwMutex() {
	// 加读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 也可以加写锁
	rwMutex.Lock()
	defer rwMutex.Unlock()

	//要执行的代码,防并发冲突
}

// 不可重入例子
func Failed1() {
	mutex.Lock()
	defer mutex.Unlock()

	// 这一句会死锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}

// 不可升级
func Failed2() {
	rwMutex.RLock()
	defer rwMutex.Unlock()

	// 这一句会死锁,读写锁里面也用了mutex锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}
