package bucket_sort

import "github.com/aikin/go-algorithms/algorithms/sorting/insertion-sort"

/*
* 其实也是分治思想

* hash function
* init bucket size
* append bucket
* sorted nums in bucket
*

*/
func Sort(nums []int) []int {
	var max, min int

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	bucketLen := int(max-min)/len(nums) + 1
	buckets := make([][]int, bucketLen)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = make([]int, 0)
	}

	for _, num := range nums {
		idx := int(num-min) / len(nums)
		buckets[idx] = append(buckets[idx], num)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if (len(bucket)) > 0 {
			insertion_sort.Sort(bucket)
			sorted = append(sorted, bucket...)
		}
	}
	return sorted
}
