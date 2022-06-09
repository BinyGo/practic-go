package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Demo1() {
	var config atomic.Value //holds current server configuration
	config.Store(loadConfig())
	//Create initial config value and store into config
	//假设一个进程对配置进行copy on write
	go func() {
		// Reload config every 10 seconds
		// and update config value with the new version
		for {
			time.Sleep(time.Second * 10)
			config.Store(loadConfig())
		}
	}()
	//另外一个进程进行读取配置,来完成无锁访问共享数据,不会race
	// Create worker goroutines that handle incoming requests
	// using the latest config value.
	for i := 0; i < 10; i++ {
		go func() {
			for n := 0; n < 5; n++ {
				c := config.Load()
				// handle request r using config c.
				fmt.Println(c)
			}
		}()
	}

}

//加载配置或者其他什么需要的数据
func loadConfig() int {
	return rand.Intn(10000-1+1) + 1
}

func Demo2() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex //used only by writers

	//read function can be used to read the data without further synchronization
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}

	//insert function can be used to update the data without further synchronization
	insert := func(key, val string) {
		mu.Lock() //synchronize with other potential writers
		defer mu.Unlock()
		m1 := m.Load().(Map) //load current value of the data structure
		m2 := make(Map)      //create a new value
		for k, v := range m1 {
			m2[k] = v // copy all data from the current object to the new one
		}
		m2[key] = val //do the update that we need
		m.Store(m2)   // atomically replace the current object with the new one
		// at this point all new readers start working with the new version.
		// The old version will be garbage collected once the existing readers
		// (is any) are done with it.
	}
	_, _ = read, insert
}
