package remove_duplicates_ii

/*
Tasking
一个排序后的数组
需求：
	* [1, 1, 1, 2, 2, 3] -> [1, 1, 1, 2, 2, 3] , output: 5
	* [0, 0, 1, 1, 1, 1, 2, 3, 3] -> [0, 0, 1, 1, 2, 3, 3], output: 7

检验：
	* 元素特点：数组内都是数字，包含整数和负数
	* 长度限制
	* 顺序不变
	* 必须是排序后数组

复杂度分析：
	* 时间复杂度：O(n)
	* 空间复杂度：O(1)

Coding
	* 中间变量：双指针，validIndexOfMaxDuplicated, loop Index
	* 遍历：一遍 for loop
	* 操作：数组内替换
		* 条件：不相等
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	countMaxDuplicated := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[countMaxDuplicated-2] {
			nums[countMaxDuplicated] = nums[i]
			countMaxDuplicated++
		}
	}
	return countMaxDuplicated
}
