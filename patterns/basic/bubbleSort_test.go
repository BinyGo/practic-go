package basic

import "testing"

//go test -v -run TestBubbleSortExec
func TestBubbleSortExec(t *testing.T) {
	BubbleSortExec()
}

func BenchmarkBubbleSortExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSortExec()
	}
}

func BenchmarkBubbleSortV2Exec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSortExecV2()
	}
}

//go test -bench=".*" -benchmem -benchtime=3s -cpu=16
/* goos: linux
goarch: amd64
pkg: github.com/practic-go/patterns/basic
cpu: 12th Gen Intel(R) Core(TM) i5-12600KF
BenchmarkBubbleSortExec-16              68395150                46.19 ns/op            0 B/op          0 allocs/op
BenchmarkBubbleSortV2Exec-16            99803300                34.82 ns/op            0 B/op          0 allocs/op
BenchmarkHeapSort-16                    47945121                74.10 ns/op            0 B/op          0 allocs/op
BenchmarkQuickSort-16                    2357053              1532 ns/op            2664 B/op         77 allocs/op
BenchmarkQuickSortV2-16                 30062385               115.7 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/practic-go/patterns/basic    19.117s */
