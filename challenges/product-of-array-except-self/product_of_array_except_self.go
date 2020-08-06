package product_of_array_except_self


/*
Given nums = [1, 2, 3, 4]

Because 24=2*3*4, 12=1*3*4 8=1*2*4 6=1*2*3
Then result: [24, 12, 8, 6]

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n)
*/
func ProductExceptSelf(nums []int) ([]int, error) {
	left, right, result := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i - 1] * nums[i - 1]
	}

	right[len(nums) - 1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i + 1] * nums[i + 1]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}

	return result, nil
}