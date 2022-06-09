package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func (c *Config) T() {}

func BenchmarkAtomic(b *testing.B) {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			//用一个新cfg对象去替换掉旧的
			cfg := &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				cfg := v.Load().(*Config)
				cfg.T()
				//fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkMutex(b *testing.B) {
	var l sync.RWMutex
	var cfg *Config

	go func() {
		i := 0
		for {
			i++
			l.Lock()
			cfg = &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			l.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				l.RLock()
				cfg.T()
				l.RUnlock()
				//fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
$ go test -bench=".*" -benchmem -benchtime=1s -cpu=2
goos: linux
goarch: amd64
pkg: github.com/practic-go/geekbang/camp/week3/atomic
cpu: 12th Gen Intel(R) Core(TM) i5-12600KF
BenchmarkAtomic-2       411036265                2.517 ns/op           2 B/op          0 allocs/op
BenchmarkMutex-2         1301199              1583 ns/op            3946 B/op        109 allocs/op
PASS
ok      github.com/practic-go/geekbang/camp/week3/atomic        4.358s
*/
