package fanOut

import (
	"fmt"
	"sync"
	"time"
)

func pool(wg *sync.WaitGroup, workers int, jobs []int) {
	//设置工作通道
	deskCh := make(chan int, workers)

	//告诉工人让他们去这个工作通道取任务
	for i := 0; i < workers; i++ {
		go worker(wg, deskCh)
	}

	//向工作通道发布任务
	for _, job := range jobs {
		deskCh <- job
	}

	//关闭工作通道（非阻塞通道一定要记得手动关闭）
	close(deskCh)
}

func worker(wg *sync.WaitGroup, deskCh <-chan int) {
	defer wg.Done()

	for {
		job, ok := <-deskCh
		if !ok {
			//任务已分配完,结束工作
			return
		}
		fmt.Println(time.Now().Unix(), "processed job ", job)
	}
}

func FanOut() {
	workerNum := 2
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 6, 6}
	var wg sync.WaitGroup
	wg.Add(workerNum)
	go pool(&wg, workerNum, jobs)
	wg.Wait()
}
