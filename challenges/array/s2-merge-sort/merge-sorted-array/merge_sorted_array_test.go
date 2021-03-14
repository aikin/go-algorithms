package merge_sorted_array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/array/s2-merge-sort/merge-sorted-array"
)

var _ = Describe("MergeSortedArray", func() {
	Context("remove element general logic", func() {
		When("nums1: [1, 2, 3, 0, 0, 0], m: 3, nums2 = [2, 5, 6], n = 3", func() {
			It("should nums1: [1, 2, 2, 3, 5, 6]", func() {
				nums1 := []int{1, 2, 3, 0, 0, 0}
				nums2 := []int{2, 5, 6}
				m := 3
				n := 3

				Merge(nums1, m, nums2, n)

				Ω(nums1).Should(Equal([]int{1, 2, 2, 3, 5, 6}))
			})

		})

		When("nums1: [1], m: 1, nums2 = [], n = 0", func() {
			It("should nums1: [1, 2, 2, 3, 5, 6]", func() {
				nums1 := []int{1}
				nums2 := []int{}
				m := 1
				n := 0

				Merge(nums1, m, nums2, n)

				Ω(nums1).Should(Equal([]int{1}))
			})
		})
	})
})
