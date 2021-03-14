package remove_element

/*
Tasking:
需求：未排序的数组，删除指定的元素
	* nums = [3, 2, 2, 3], val = 3; => [2, 2] , output: 2
	* nums = [0, 1, 2, 2, 3, 0, 4, 2], val = 2; => [0, 1, 3, 0, 4], output: 5

检验：
	* 数组内都是数字，包含整数和负数
	* 指定元素为数字
	* 长度限制
	* 顺序：可以任意顺序

复杂度分析：
	* 时间复杂度：O(n)
	* 空间复杂度：O(1)

Coding
	* 中间变量：双指针
	* 遍历：一遍 for loop
	* 操作：swap
*/
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	return slow
}
