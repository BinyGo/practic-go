package basic

// 冒泡排序算法
// 算法描述：冒泡算法，数组中前一个元素和后一个元素进行比较如果大于或者小于 前者就进行交换，最终返回最大或者最小都冒到数组的最后序列时间复杂度为 O(n^2).

// 算法步骤
// 从数组开头选择相邻两个元素进行比较，并进行交换
// 不停向后移动

func BubbleSortExec() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	BubbleSort(arr)
	//fmt.Println(BubbleSort(arr))

}

func BubbleSortExecV2() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	BubbleSortV2(arr)
	//fmt.Println(BubbleSortV2(arr))
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func BubbleSortV2(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	for i := 0; i < n; i++ {
		is_sorted := true
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				is_sorted = false
			}
		}
		if is_sorted {
			break
		}
	}
	return arr
}

func GetMax(arr []int) int {
	for j := 1; j < len(arr); j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr[len(arr)-1]
}
