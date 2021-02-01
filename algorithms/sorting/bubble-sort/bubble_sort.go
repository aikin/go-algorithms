package bubble_sort

func Sort(nums []int) {
	for unsortedLen := len(nums) - 1; ; unsortedLen-- {
		swap := false
		for i := 0; i < unsortedLen; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}
