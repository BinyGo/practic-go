package basic

import "testing"

func TestExecHeapSort(t *testing.T) {
	HeapSortExec()
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSortExec()
	}
}

/*
go test -bench=".*" -benchmem -benchtime=3s -cpu=16
goos: linux
goarch: amd64
pkg: github.com/practic-go/patterns/basic
cpu: 12th Gen Intel(R) Core(TM) i5-12600KF
BenchmarkHeapSort-16            47231149                73.42 ns/op            0 B/op          0 allocs/op
BenchmarkQuickSort-16            2358675              1528 ns/op            2664 B/op         77 allocs/op
BenchmarkQuick-16               31276843               114.9 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/practic-go/patterns/basic    12.411s
*/
