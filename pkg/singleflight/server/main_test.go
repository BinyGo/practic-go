package main

import (
	"sync"
	"testing"
)

func TestDemo1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func() {
			defer wg.Done()
			Client()
		}()
	}
	wg.Wait()

	//time.Sleep(time.Second * 1)
}
