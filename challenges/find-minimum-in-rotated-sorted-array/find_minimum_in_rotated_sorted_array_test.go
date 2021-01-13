package find_minimum_in_rotated_sorted_array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/find-minimum-in-rotated-sorted-array"
)

var _ = Describe("FindMinimumInRotatedSortedArray", func() {
	Context("find min general logic", func() {
		When("numbres [3, 4, 5, 1, 2]", func() {
			It("should return min num: 1", func() {

				var nums = []int{3, 4, 5, 1, 2}

				min := FindMin(nums)

				Ω(min).Should(Equal(1))
			})
		})
	})
})
