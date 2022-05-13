package basic

// 算法描述：从未排序数据中选择最大或者最小的值和当前值交换 O(n^2).

// 算法步骤
// 选择一个数当最小值或者最大值，进行比较然后交换
// 循环向后查进行
func SelectSortExec() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	SelectSort(arr)
	//selectsort := SelectSort(arr)
	//fmt.Println(selectsort)
}

//切片排序
func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

//获取切片里面的最大值
func SelectMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	}
	max := arr[0]
	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//go test -bench=".*" -benchmem -benchtime=3s -cpu=16
/*
goos: linux
goarch: amd64
pkg: github.com/practic-go/patterns/basic
cpu: 12th Gen Intel(R) Core(TM) i5-12600KF
BenchmarkBinSearchExec-16               1000000000               3.297 ns/op           0 B/op          0 allocs/op
BenchmarkBubbleSortExec-16              99731878                34.92 ns/op            0 B/op          0 allocs/op
BenchmarkBubbleSortV2Exec-16            100000000               33.70 ns/op            0 B/op          0 allocs/op
BenchmarkHeapSort-16                    48979960                72.48 ns/op            0 B/op          0 allocs/op
BenchmarkQuickSort-16                    2380426              1527 ns/op            2664 B/op         77 allocs/op
BenchmarkQuickSortV2-16                 31408530               115.3 ns/op             0 B/op          0 allocs/op
BenchmarkSelectSort-16                   2362952              1522 ns/op            2664 B/op         77 allocs/op
PASS
ok      github.com/practic-go/patterns/basic    28.239s
*/
