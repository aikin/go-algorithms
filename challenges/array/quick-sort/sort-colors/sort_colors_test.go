package sort_colors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/array/quick-sort/sort-colors"
)

var _ = Describe("SortColors", func() {
	Context("sort colors general logic", func() {
		When("when input nums is : [2, 0, 2, 1, 1, 0]", func() {
			It("should return length: [0, 0, 1, 1, 2, 2]", func() {

				var nums = []int{2, 0, 2, 1, 1, 0}

				SortColors(nums)

				Ω(nums).Should(Equal([]int{0, 0, 1, 1, 2, 2}))
			})
		})
	})
})
