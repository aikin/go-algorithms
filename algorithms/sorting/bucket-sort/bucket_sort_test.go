package bucket_sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/algorithms/sorting/bucket-sort"
)

var _ = Describe("BucketSort", func() {
	Context("When sort nums: [8, 3, 4, 5, 6, 7, 1, 9, 2, 0]", func() {
		It("should sorted nums: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]", func() {
			nums := []int{8, 3, 4, 5, 6, 7, 1, 9, 2, 0}

			sorted := Sort(nums)

			Ω(sorted).Should(Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})
	})

	Context("When sort nums: [1, 0, 5, 10, 20, 13, 7, 3, 2, 3]", func() {
		It("should sorted nums: [0, 1, 2, 3, 3, 5, 7, 10, 13, 20]", func() {
			nums := []int{1, 0, 5, 10, 20, 13, 7, 3, 2, 3}

			sorted := Sort(nums)

			Ω(sorted).Should(Equal([]int{0, 1, 2, 3, 3, 5, 7, 10, 13, 20}))
		})
	})
})
