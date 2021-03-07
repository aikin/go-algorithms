package counting_sort

func Sort(nums []int, maxValue int) {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	sortedIndex := 0

	for i := 0; i < len(nums); i++ {
		bucket[nums[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			nums[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
}
