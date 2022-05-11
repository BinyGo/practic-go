package basic

import (
	"reflect"
	"testing"
)

//go test -v -run TestQuickSort
func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one", args{arr: []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}}, []int{1, 2, 5, 8, 9, 10, 12, 30, 45, 63, 234}},
		{"two", args{arr: []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}}, []int{1, 2, 5, 8, 9, 10, 12, 30, 45, 63, 234}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_Quick(t *testing.T) {
	QuickSortExec()
	QuickSortExecV2()
}

func BenchmarkQuickSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		QuickSortExec()
	}
}

func BenchmarkQuickSortV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSortExecV2()
	}
}

// go test -bench=".*" -benchmem -benchtime=10s -cpu=10  //-benchmem #输出内存分配统计  -benchtime=10s # 指定测试时间
// goos: linux
// goarch: amd64
// pkg: github.com/practic-go/patterns/basic
// cpu: 12th Gen Intel(R) Core(TM) i5-12600KF
// BenchmarkQuickSort-16            1414204               847.9 ns/op          1216 B/op         50 allocs/op
// BenchmarkQuick-16               23924130                48.88 ns/op            0 B/op          0 allocs/op
// PASS
// ok      github.com/practic-go/patterns/basic    3.283s

//BenchmarkQuickSort-16，BenchmarkQuickSort表示所测试的测试函数名，16 表示有 16 个 CPU 线程参与了此次测试，默认是GOMAXPROCS的值。
//1414204 ，说明函数中的循环执行了90848414次。
//847.9 ns/op 说明每次循环的执行平均耗时是 847.9 纳秒，该值越小，说明代码性能越高。
//1216 B/op  表示每次执行分配了多少内存（字节），该值越小，说明代码内存占用越小；
//50 allocs/op 表示每次执行分配了多少次内存，该值越小，说明分配内存次数越少，意味着代码性能越高。

//秒的换算:ms(毫秒),μs(微秒),ns(纳秒),ps(皮秒)
