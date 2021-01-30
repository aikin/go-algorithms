package merge_sorted_array

import "sort"

/*
Tasking:
需求：Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

检验：
	* nums1.length = m+n; nums2.length = n
	* 0 <= m, n <= 200
	* 1 <= m + n <= 200
	* -109 <= nums1[i], nums2[i] <= 109
	* 数组内为有序整数

复杂度分析：
	* 时间复杂度：O(m+n)
	* 空间复杂度：O(m+n)

Coding
	* 中间变量：双指针
	* 遍历：一遍 for loop
	* 操作：swap
*/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	aux := []int{}
	i := 0
	j := 0

	for len(aux) < m+n {
		if i >= m {
			aux = append(aux, nums2[j])
			j++
		} else if j >= n {
			aux = append(aux, nums1[i])
			i++
		} else if nums1[i] <= nums2[j] {
			aux = append(aux, nums1[i])
			i++
		} else {
			aux = append(aux, nums2[j])
			j++
		}
	}

	for l := 0; l < len(aux); l++ {
		nums1[l] = aux[l]
	}
}

func MergeBySort(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(append(nums1[:m], nums2...))
}
