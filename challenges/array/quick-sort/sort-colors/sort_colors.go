package sort_colors

/*
Tasking:
需求：
	* [2, 0, 2, 1, 1, 0] ->  output: [0, 0, 1, 1, 2, 2]
	* 只会出现 0, 1, 2
	* 必须在原数组上操作，不能拷贝额外的数组。

检验：
	* 数组内只能出现 0, 1, 2
	* 长度限制 1 <= n <= 300

复杂度分析：
	* 时间复杂度：O(n)
	* 空间复杂度：O(1)

Coding
	* 中间变量：单指针，替换为。
	* 遍历：一遍 for loop
	* 操作：数组内替换
*/
func SortColors(nums []int) {
	n := len(nums)
	if n > 300 {
		return
	}

	countZero := swapColors(nums, 0)
	swapColors(nums[countZero:], 1)
}

func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}
