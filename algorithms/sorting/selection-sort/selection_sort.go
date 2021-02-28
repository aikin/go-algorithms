package selection_sort

func Sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		maxIndex := 0
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[len(nums)-i-1], nums[maxIndex] = nums[maxIndex], nums[len(nums)-i-1]
	}
}
