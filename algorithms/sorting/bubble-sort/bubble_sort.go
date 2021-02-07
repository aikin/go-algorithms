package bubble_sort

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
	for unsortedLen := len(nums) - 1; ; unsortedLen-- {
		swap := false
		for i := 0; i < unsortedLen; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}
