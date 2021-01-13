package find_minimum_in_rotated_sorted_array

/*
Given: nums = [3,4,5,1,2]
WHEN: The original array was [1,2,3,4,5] rotated 3 times.
THEN: 1

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n)
*/
func FindMinByForceWay(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if (nums[i] < min) {
			min = nums[i]
		}
	}
	return min
}