package merge_sort

func Sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums) / 2
	left := Sort(nums[:middle])
	right := Sort(nums[middle:])

	return mergeSortedNums(left, right)
}

func mergeSortedNums(leftSortedNums, rightSortedNums []int) []int {
	sortedNums := make([]int, len(leftSortedNums)+len(rightSortedNums))

	for i := 0; len(leftSortedNums) > 0 || len(rightSortedNums) > 0; i++ {
		if len(leftSortedNums) > 0 && len(rightSortedNums) > 0 {
			if leftSortedNums[0] < rightSortedNums[0] {
				sortedNums[i] = leftSortedNums[0]
				leftSortedNums = leftSortedNums[1:]
			} else {
				sortedNums[i] = rightSortedNums[0]
				rightSortedNums = rightSortedNums[1:]
			}

		} else if len(leftSortedNums) > 0 {

			sortedNums[i] = leftSortedNums[0]
			leftSortedNums = leftSortedNums[1:]

		} else if len(rightSortedNums) > 0 {
			sortedNums[i] = rightSortedNums[0]
			rightSortedNums = rightSortedNums[1:]
		}
	}
	return sortedNums
}
