package product_of_array_except_self

import (
	"errors"
)

const uintSize = 32 << (^uint(0) >> 32 & 1)
const (
	maxInt  = 1<<(uintSize-1) - 1
	minInt  = -maxInt - 1
	maxUint = 1<<uintSize - 1
)

/*
Given nums = [1, 2, 3, 4]

Because 24=2*3*4, 12=1*3*4 8=1*2*4 6=1*2*3
Then result: [24, 12, 8, 6]

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n)
*/
func ProductExceptSelf(nums []int) ([]int, error) {
	left, right, products := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		if isNumLesserThanOne(nums[i]) {
			return nil, errors.New("the num > 1 is required")
		}
		if isProductOversize(left[i-1], nums[i-1]) {
			return nil, errors.New("the num should not oversize int")
		}
		left[i] = left[i-1] * nums[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		if isNumLesserThanOne(nums[i]) {
			return nil, errors.New("the num > 1 is required")
		}
		if isProductOversize(right[i+1], nums[i+1]) {
			return nil, errors.New("the num should not oversize int")
		}
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		if isProductOversize(left[i], right[i]) {
			return nil, errors.New("the num should not oversize int")
		}
		products[i] = left[i] * right[i]
	}

	return products, nil
}

func isProductOversize(multiplier int, multiplicand int) bool {
	return multiplier * multiplicand < multiplicand && (multiplicand > 1)
}

func isNumLesserThanOne(factor int) bool {
	return factor < 1
}
