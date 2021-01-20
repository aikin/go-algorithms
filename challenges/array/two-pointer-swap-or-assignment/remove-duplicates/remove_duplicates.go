package remove_duplicates

/*
Tasking:
需求：排序后的
	* [1, 1, 2] -> [1, 2] , output: 2
	* [1] -> [1], output: 1
	* [0, 0, 1, 1, 1, 2, 2, 3, 3, 4] -> [0, 1, 2, 3, 4], output: 4
	* [-1, -2, 1, 1] -> [-1, -2, 1], output: 3

检验：
	* 数组内都是数字，包含整数和负数
	* 长度限制
	* 顺序不变

复杂度分析：
	* 时间复杂度：O(n)
	* 空间复杂度：O(1)

Coding
	* 中间变量：
	* 遍历：一遍 for loop
	* 操作：数组内替换
*/
func RemoveDuplicates (nums []int) int {
	if (len(nums) <= 1) {
		return len(nums)
	}

	countUniqueNum := 1

	for i := 0; i < len(nums) - 1; i++ {
		if (nums[i] != nums[i + 1]) {
			countUniqueNum++
		}
	}

	return countUniqueNum
}
