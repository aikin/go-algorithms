package selection_sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/algorithms/sorting/selection-sort"
)

var _ = Describe("Selection Sort", func() {
	Context("When sort nums: [8, 3, 4, 5, 6, 7, 1, 9, 2, 0]", func() {
		It("should sorted nums: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]", func() {
			nums := []int{8, 3, 4, 5, 6, 7, 1, 9, 2, 0}

			Sort(nums)

			Ω(nums).Should(Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})
	})

	Context("When sort nums: [-1, 0, 5, -10, 20, 13, -7, 3, 2, -3]", func() {
		It("should sorted nums: [-10, -7, -3, -1, 0, 2, 3, 5, 13, 20]", func() {
			nums := []int{-1, 0, 5, -10, 20, 13, -7, 3, 2, -3}

			Sort(nums)

			Ω(nums).Should(Equal([]int{-10, -7, -3, -1, 0, 2, 3, 5, 13, 20}))
		})
	})
})
