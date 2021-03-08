package kth_largest_element_in_an_array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/aikin/go-algorithms/challenges/array/quick-sort/kth-largest-element-in-an-array"
)

var _ = Describe("KthLargestElementInAnArray", func() {
	Context("kth largest element in an array general logic", func() {
		When("when input nums is : [3, 2, 1, 5, 6, 4]", func() {
			It("should return kth largest: 2", func() {

				var nums = []int{3, 2, 1, 5, 6, 4}

				kthLargest := FindKthLargestWithMinHeap(nums, 2)

				Ω(kthLargest).Should(Equal(5))
			})
		})

		When("when input nums is : [3, 2, 3, 1, 2, 4, 5, 5, 6]", func() {
			It("should return kth largest: 4", func() {

				var nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}

				kthLargest := FindKthLargestWithMinHeap(nums, 4)

				Ω(kthLargest).Should(Equal(4))
			})
		})
	})
})
