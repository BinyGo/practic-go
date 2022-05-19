package main

import (
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	for i := 0; i < 3; i++ {
		Client()
	}
	Client2()
	time.Sleep(time.Second * 1)
}
