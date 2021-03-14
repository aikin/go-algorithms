package move_zeroes

/*
Tasking:
需求：
	* [0, 1, 0, 3, 12] ->  output: [1, 3, 12, 0, 0]
	* [1, 3, 2, 0, 0, 1, 3, 4] -> output: [1, 3, 2, 1, 3, 4, 0, 0]
	* 必须在原数组上操作，不能拷贝额外的数组。
	* 编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

检验：
	* 数组内都是数字，包含整数和负数
	* 长度限制
	* 顺序不变

复杂度分析：
	* 时间复杂度：O(n)
	* 空间复杂度：O(1)

Coding
	* 中间变量：双指针
	* 遍历：一遍 for loop
	* 操作：数组内替换
*/

func MoveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
