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
 - 时间复杂度：O(n^3)
 - 空间复杂度：O(1)
*/
func ThreeSum(nums []int) [][]int {

	ans := make([][]int, 0)
	sort.Ints(nums);
	l := len(nums)

	for first := 0; first < l; first++ {
		if (first > 0 && nums[first] == nums[first - 1]) {
			continue
		}

		third := l - 1
		target := -1 * nums[first]

		for second := first + 1; second < l; second++ {
			if (second > first + 1 && nums[second] == nums[second -1]) {
				continue
			}

			if second < third && nums[second] + nums[third] > target {
				third--
			}

			if (second == third) {
				break
			}

			if (nums[second] + nums[third] == target) {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}

	}

	return ans
}