package two_sum

import (
	// "math"
	// "errors"
	"fmt"
)

/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func TwoSum(nums []int, target int) []int {
	fmt.Printf("starting")

	var indices []int
	var exceptedIndicesSize = 2

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		for j := 0; j < len(nums); j++ {
			if (cur + nums[j] == target) {
				indices = append(indices, i, j)
				break
			}
		}
		if (len(indices) == exceptedIndicesSize) {
			break
		}
	}

	return indices
}