package two_sum

import (
	// "math"
	"errors"
	// "fmt"
	// "strconv"
)

const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
const (
	maxInt  = 1<<(uintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	minInt  = -maxInt - 1         // -1 << 31 or -1 << 63
	maxUint = 1<<uintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func TwoSum(nums []int, target int) ([]int, error) {
	var indices []int
	exceptedIndicesSize := 2


	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			hasDuplicated := nums[i] == nums[j]
			if (hasDuplicated) {
				return nil, errors.New("nums duplicates")
			}

			isOverflowSize := nums[i] + nums[j] <= nums[i] == (nums[j] > 0)
			/*
			// isOverflowSize := (nums[i] + nums[j] <= nums[i] && nums[j] > 0) ||
			// (nums[i] + nums[j] >= nums[i] && nums[j] < 0)
			*/
			if (isOverflowSize)  {
				return nil, errors.New("should two sum in [min, max] rang")
			}

			isNotFull := len(indices) != exceptedIndicesSize
			isTwoSumEqualTarget := nums[i] + nums[j] == target

			if (isNotFull && isTwoSumEqualTarget) {
				indices = append(indices, i, j)
			}
			if (!isNotFull && isTwoSumEqualTarget) {
				return nil, errors.New("should exactly one solution")
			}
		}
	}

	return indices, nil
}