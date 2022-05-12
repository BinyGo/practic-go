# patterns basic  

单元测试  go test -v -run TestBinSearchExec

基准测试 go test -bench=".*" -benchmem -benchtime=3s -cpu=16
<!-- 
goos: linux
goarch: amd64
pkg: github.com/practic-go/patterns/basic
cpu: 12th Gen Intel(R) Core(TM) i5-12600KF
BenchmarkBinSearchExec-16               1000000000               3.327 ns/op           0 B/op          0 allocs/op
BenchmarkBubbleSortExec-16              83262081                44.12 ns/op            0 B/op          0 allocs/op
BenchmarkBubbleSortV2Exec-16            100000000               35.80 ns/op            0 B/op          0 allocs/op
BenchmarkHeapSort-16                    47900617                75.02 ns/op            0 B/op          0 allocs/op
BenchmarkQuickSort-16                    2376908              1519 ns/op            2664 B/op         77 allocs/op
BenchmarkQuickSortV2-16                 30960536               113.6 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/practic-go/patterns/basic    23.461s -->