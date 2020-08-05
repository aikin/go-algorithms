package maximum_subarray_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/maximum-subarray"
)

var _ = Describe("MaximumSubarray", func() {
	Context("Maximum Subarray general logic", func() {
		When("nums [-2, 1, -3, 4, -1, 2, 1, -5, 4]", func() {
			It("should return maxSum: 6, err: nil", func() {

				nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

				maxSum, err := MaxSubArrayByBruteForceWay(nums)

				Ω(err).Should(BeNil())
				Ω(maxSum).Should(Equal(6))
			})
		})

		When("nums [1]", func() {
			It("should return maxSum: 1, err: nil", func() {

				nums := []int{1}

				maxSum, err := MaxSubArrayByBruteForceWay(nums)

				Ω(err).Should(BeNil())
				Ω(maxSum).Should(Equal(1))
			})
		})

		When("nums [-1]", func() {
			It("should return maxSum: -1, err: nil", func() {

				nums := []int{-1}

				maxSum, err := MaxSubArrayByBruteForceWay(nums)

				Ω(err).Should(BeNil())
				Ω(maxSum).Should(Equal(-1))
			})
		})
	})
})
