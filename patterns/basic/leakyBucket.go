package basic

import (
	"fmt"
	"sync"
	"time"
)

// 漏桶算法
// 适用场景
// 流量最均匀的限流方式，一般用于流量“整形”，例如保护数据库的限流。先把对数据库的访问加入到木桶中，
// worker再以db能够承受的qps从木桶中取出请求，去访问数据库。不太适合电商抢购和微博出现热点事件等场景的限流，
// 一是应对突发流量不是很灵活，二是为每个user_id/ip维护一个队列(木桶)，workder从这些队列中拉取任务，资源的消耗会比较大。

// 每个请求来了，把需要执行的业务逻辑封装成Task，放入木桶，等待worker取出执行
type Task struct {
	handler func() Result // worker从木桶去除请求对象后要执行的业务逻辑函数
	resChan chan Result   // 等待worker执行并返回结果的channel
	taskID  int
}

type Result struct{}

//模拟业务逻辑函数
func handler() Result {
	time.Sleep(time.Millisecond * 300)
	return Result{}
}

func NewTask(id int) Task {
	return Task{
		handler: handler,
		resChan: make(chan Result),
		taskID:  id,
	}
}

//漏桶
type LeakyBucket struct {
	BucketSize int       // 木桶大小
	NumWorker  int       // 同时从木桶中获取任务执行的worker数量
	bucket     chan Task //存放任务的木桶
}

func NewLeakyBucket(bucketSize int, numWorker int) *LeakyBucket {
	return &LeakyBucket{
		BucketSize: bucketSize,
		NumWorker:  numWorker,
		bucket:     make(chan Task, bucketSize),
	}
}

func (b *LeakyBucket) Validate(task Task) bool {
	//如果木桶已经满了,返回false
	select {
	case b.bucket <- task:
	default:
		fmt.Printf("request[id=%d] is refused\n", task.taskID)
		return false
	}
	//等待worker执行
	<-task.resChan
	fmt.Printf("request[id=%d] is run\n", task.taskID)
	return true
}

func (b *LeakyBucket) Start() {
	//开启worker从木桶拉取任务执行
	go func() {
		for i := 0; i < b.NumWorker; i++ {
			go func() {
				task := <-b.bucket
				result := task.handler()
				task.resChan <- result
			}()
		}
	}()
}

func LeakyBucketExec() {
	bucket := NewLeakyBucket(10, 4)
	bucket.Start()
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := NewTask(id)
			bucket.Validate(task)
		}(i)
	}
	wg.Wait()
}
