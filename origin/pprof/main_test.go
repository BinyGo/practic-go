package main

import (
	_ "net/http/pprof"
	"testing"
)

func TestAdd(t *testing.T) {
	_ = Add("go-biny-test")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add("go-biny-test")
	}
}

//$ go test -bench=. -cpuprofile=cpu.profile
// 生成 cpu.profile 文件，然后执行  go tool pprof cpu.profile 命令查看

//$ go test -bench=. -memprofile=mem.profile
// 生成 mem.profile 文件，然后执行  go tool pprof mem.profile 命令查看
