package bubble_sort_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/algorithms/sorting/bubble-sort"
)

var _ = Describe("BubbleSort", func() {
	Context("When sort nums: [8, 3, 4, 5, 6, 7, 1, 9, 2, 0]", func() {
		It("should sorted nums: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]", func() {
			nums := []int{8, 3, 4, 5, 6, 7, 1, 9, 2, 0}

			Sort(nums)

			Ω(nums).Should(Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})
	})
})
