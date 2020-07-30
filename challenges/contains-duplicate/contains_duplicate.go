package contains_duplicate

import (
	"sort"
)

/*
Given nums = [1, 2, 3, 1]

Because nums[0] - nums[3] = 0
return false
复杂度分析：
 - 时间复杂度：O(n^2)
 - 空间复杂度：O(1)，只需要常数级别的存储
*/
func ContainsDuplicate(nums []int) (bool, error) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] == nums [j]) {
				return true, nil
			}
		}
	}

	return false, nil
}

/*
Given nums = [1, 2, 3, 1]

Because nums[0] - nums[3] = 0
return false
复杂度分析：
 - 时间复杂度：O(nlogn) 排序的复杂度是 O(nlogn)，扫描的复杂度是 O(n)。整个算法主要由排序过程决定，因此是 O(nlogn)
 - 空间复杂度：O(1)，只需要常数级别的存储
*/
func ContainsDuplicateAfterSort(nums []int) (bool, error) {
 	sort.Ints(nums)
	for i := 0; i < len(nums) - 1; i++ {
			if (nums[i] == nums [i + 1]) {
				return true, nil
			}
	}
	return false, nil
}

/*
Given nums = [1, 2, 3, 1]

Because nums[0] - nums[3] = 0
return false
复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(n)
*/
func ContainsDuplicateWithHashMap(nums []int) (bool, error) {
	numsMap := make(map[int]int)

	for i, curNum := range nums {
	 if _, ok := numsMap[curNum]; ok {
		return true, nil
	 }
	 numsMap[curNum] = i
 }
 return false, nil
}



