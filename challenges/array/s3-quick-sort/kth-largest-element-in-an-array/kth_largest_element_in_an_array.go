package kth_largest_element_in_an_array

import (
	. "github.com/aikin/go-algorithms/data-structures/heap"
)

/*
Tasking:
需求：
	* [3, 2, 1, 5, 6, 4], 2 ->  output: 5
	* 未排序的数组
	* 第 K 个最大元素

检验：
	* 数组内都是数字，包含整数和负数
	* 长度限制
	* 1<= k <= len(nums)

复杂度分析：
	* 时间复杂度：O(nlogk)
	* 空间复杂度：O(k)

Coding
	* 中间变量：
		- 堆数据
	* 遍历：一遍 for loop
	* 操作：遍历填充进堆数据， 等于 K 时，取出数据
*/
func FindKthLargestWithMinHeap(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1
	}
	heap := NewMinHeap()

	for _, num := range nums {
		heap.Insert(Int(num))
		if heap.Len() == k+1 {
			heap.Poll()
		}
	}
	return int(heap.Poll().(Int))
}

/*
复杂度分析：
* 时间复杂度：O(n)
* 空间复杂度：O(1)

Coding
* 中间变量：
- 堆数据
* 遍历：一遍 for loop
* 操作：遍历填充进堆数据， 等于 K 时，取出数据
*/
func FindKthLargestWithQuickSelect(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return -1
	}
	return 1
}
