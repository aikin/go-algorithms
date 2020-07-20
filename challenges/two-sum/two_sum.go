package two_sum

import (
	// "math"
	"errors"
	"fmt"
)

/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func TwoSum(nums []int, target int) ([]int, error) {
	fmt.Printf("starting")


	var indices []int
	exceptedIndicesSize := 2

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			hasDuplicated := nums[i] == nums[j]
			if (hasDuplicated) {
				return nil, errors.New("nums duplicates")
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