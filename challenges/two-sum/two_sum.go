package two_sum

import (
	// "math"
	"errors"
	// "fmt"
	// "strconv"
)

const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
const (
	maxInt  = 1 << (uintSize - 1) - 1 // 1<<31 - 1 or 1<<63 - 1
	minInt  = -maxInt - 1         // -1 << 31 or -1 << 63
	maxUint = 1 << uintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

复杂度分析：
 - 时间复杂度：O(n)：我们只遍历了包含有 nn 个元素的列表一次。在表中进行的每次查找只花费 O(1)O(1) 的时间。
 - 空间复杂度：O(n)：所需的额外空间取决于哈希表中存储的元素数量，该表最多需要存储 nn 个元素。
*/
func TwoSum(nums []int, target int) ([]int, error) {

	numsMap := make(map[int]int)  // Num as key, Index as Value
	indices := []int{}
	exceptedIndicesSize := 2

	for i, curNum := range nums {
		_, existed := numsMap[curNum]
		if (existed) {
			return nil, errors.New("nums duplicates")
		}

		isOverflowSize := target - curNum >= target == (curNum > 0)
		/*
			// isOverflowSize := (nums[i] + nums[j] <= nums[i] && nums[j] > 0) ||
			// (nums[i] + nums[j] >= nums[i] && nums[j] < 0)
			*/
		if (isOverflowSize)  {
			return nil, errors.New("should two sum in [min, max] rang")
		}

		isNotFull := len(indices) != exceptedIndicesSize
		j, existed := numsMap[target - curNum]
		if (isNotFull && existed) {
			indices = append(indices, j, i)
		}

		if (!isNotFull && existed) {
			return nil, errors.New("should exactly one solution")
		}

		numsMap[curNum] = i
	}
	return indices, nil
}