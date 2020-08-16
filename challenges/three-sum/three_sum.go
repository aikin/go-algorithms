package three_sum

import (
	"sort"
)

/*
Given nums = [-1, 0, 1, 2, -1, -4]

Because a + b + c = 0
Then result:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

复杂度分析：
 - 时间复杂度：O(n^2)
 - 空间复杂度：O(logN)
*/
func ThreeSum(nums []int) [][]int {

	ans := make([][]int, 0)
	sort.Ints(nums);
	l := len(nums)

	for i := 0; i < l; i++ {
		if (i > 0 && nums[i] == nums[i - 1]) {
			continue
		}

		right := l - 1
		target := -1 * nums[i]

		for left := i + 1; left < l; left++ {
			if (left > i + 1 && nums[left] == nums[left -1]) {
				continue
			}

			if left < right && nums[left] + nums[right] > target {
				right--
			}

			if (left == right) {
				break
			}

			if (nums[left] + nums[right] == target) {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
			}
		}

	}

	return ans
}