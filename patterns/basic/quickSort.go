package basic

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]          //第一个数据
	low := make([]int, 0)        //比我小的数据
	hight := make([]int, 0)      //比我大的数据
	mid := make([]int, 0)        //一样大的数据
	mid = append(mid, splitData) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myArr := append(append(low, mid...), hight...)
	return myArr
}

func QuickSortExec() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12, 1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	QuickSort(arr)
	//fmt.Println(QuickSort(arr))

}

func QuickSortExecV2() {
	arr2 := []int{1, 9, 11, 30, 2, 5, 45, 8, 63, 235, 12, 1, 9, 11, 30, 2, 5, 45, 8, 63, 235, 12}
	Quick(arr2, 0, len(arr2)-1) //原地排序,最优解法
	//fmt.Println(arr2)
}

//原地排序,利用了左右指针.双边循环交换法,基准测试是QuickSort的10几倍,且不需额外内存,原地交换排序
func Quick(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	Quick(arr, start, pivot-1)
	Quick(arr, pivot+1, end)

}

func partition(arr []int, start, end int) int {
	pivot := arr[start]
	left := start
	right := end
	for left < right {
		for left < right && arr[right] > pivot {
			right -= 1
		}
		for left < right && arr[left] <= pivot {
			left += 1
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[start] = arr[left]
	arr[left] = pivot
	return left
}
