package radix_sort

func findLargestNum(nums []int) int {
	largestNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > largestNum {
			largestNum = nums[i]
		}
	}
	return largestNum
}

func Sort(nums []int) {
	largestNum := findLargestNum(nums)

	size := len(nums)

	significantDigit := 1
	semiSorted := make([]int, size, size)

	for largestNum/significantDigit > 0 {
		bucket := [10]int{0}

		for i := 0; i < size; i++ {
			bucket[(nums[i]/significantDigit)%10]++
		}

		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := size - 1; i >= 0; i-- {
			bucket[(nums[i]/significantDigit)%10]--
			semiSorted[bucket[(nums[i]/significantDigit)%10]] = nums[i]
		}

		for i := 0; i < size; i++ {
			nums[i] = semiSorted[i]
		}
		significantDigit += 10
	}

}
