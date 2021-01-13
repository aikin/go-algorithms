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

				// min := FindMinByForceWay(nums)
				min := FindMinByBinarySearch(nums)

				Ω(min).Should(Equal(1))
			})
		})

		When("numbres [4,5,6,7,0,1,2]", func() {
			It("should return min num: 0", func() {

				var nums = []int{4, 5, 6, 7, 0, 1, 2}

				// min := FindMinByForceWay(nums)
				min := FindMinByBinarySearch(nums)

				Ω(min).Should(Equal(0))
			})
		})

		When("numbres [11,13,15,17]", func() {
			It("should return min num: 11", func() {

				var nums = []int{11, 13, 15, 17}

				// min := FindMinByForceWay(nums)
				min := FindMinByBinarySearch(nums)

				Ω(min).Should(Equal(11))
			})
		})

		When("numbres [3]", func() {
			It("should return min num: 1", func() {

				var nums = []int{3}

				// min := FindMinByForceWay(nums)
				min := FindMinByBinarySearch(nums)

				Ω(min).Should(Equal(3))
			})
		})
	})
})
