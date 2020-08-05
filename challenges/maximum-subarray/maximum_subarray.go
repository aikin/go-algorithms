package maximum_subarray

const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
const (
	maxInt  = 1<<(uintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	minInt  = -maxInt - 1         // -1 << 31 or -1 << 63
	maxUint = 1<<uintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

/*
Given nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]

Because contiguous subarray [4,-1,2,1] has the largest sum = 6
Then max = 6

复杂度分析：
 - 时间复杂度：O(n^2)
 - 空间复杂度：O(1)
*/
func MaxSubArrayByBruteForceWay(nums []int) (int, error) {
	max := minInt

	for i := 0; i < len(nums); i++ {
		curSum := 0
		for j := i; j < len(nums); j++ {
			curSum += nums[j]
			if curSum > max {
				max = curSum
			}
		}
	}

	return max, nil
}

/*
Given nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]

Because contiguous subarray [4,-1,2,1] has the largest sum = 6
Then max = 6

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(1)
*/
func MaxSubArrayByGreedyWay(nums []int) (int, error) {
	max := minInt
	curSum := 0

	for _, curNum := range nums {
		curSum += curNum
		if (curSum > max) {
			max = curSum
		}
		if (curSum < 0) {
			curSum = 0
		}
	}
	return max, nil
}

/*
Given nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]

Because contiguous subarray [4,-1,2,1] has the largest sum = 6
Then max = 6

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n)
*/
func MaxSubArrayByDynamicWay(nums []int) (int, error) {
	sums := make([]int, len(nums))
	sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if (sums[i - 1] < 0) {
			sums[i] = nums[i]
			continue
		}
		sums[i] = sums[i - 1] + nums[i]
	}

	max := minInt
	for _, curSum := range sums {
		if (curSum > max) {
			max = curSum
		}
	}

	return max, nil
}