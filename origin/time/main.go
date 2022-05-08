package main

import (
	"fmt"
	"time"
)

//timer定时器
func Timer() {
	for {
		now := time.Now()
		//1分钟后
		mm, _ := time.ParseDuration("10s") //10m //1h
		next := now.Add(mm)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		fmt.Println("start Timer task")
	}

}
