package quick_sort

/*
复杂度分析：
	* 时间复杂度：O(n2)
	* 空间复杂度：O(1)

Coding
	* 中间变量：双指针
	* 遍历：twice for loop
	* 操作：如果左边>右边，就 swap。
*/
func Sort(nums []int) {
	recursionSort(nums, 0, len(nums)-1)
}

func recursionSort(data []int, left int, right int) {
	if left < right {
		pivot := partition(data, left, right)
		recursionSort(data, left, pivot-1)
		recursionSort(data, pivot+1, right)
	}
}

func partition(data []int, left int, right int) int {
	for left < right {
		for left < right && data[left] <= data[right] {
			right--
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			left++
		}

		for left < right && data[left] <= data[right] {
			left++
		}

		if left < right {
			data[left], data[right] = data[right], data[left]
			right--
		}
	}
	return left
}
