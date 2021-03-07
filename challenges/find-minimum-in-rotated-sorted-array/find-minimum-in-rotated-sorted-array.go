package find_minimum_in_rotated_sorted_array

/*
Given: nums = [3,4,5,1,2]
WHEN: The original array was [1,2,3,4,5] rotated 3 times.
THEN: 1

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(1)
*/
func FindMinByBruteForceWay(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

/*
复杂度分析：
 - 时间复杂度：O(logn)
 - 空间复杂度：O(1)
*/
func FindMinByBinarySearch(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	if nums[right] > nums[0] {
		return nums[0]
	}

	for left <= right {

		if nums[left] <= nums[right] {
			return nums[left]
		}

		mid := left + (right-left)>>1
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}