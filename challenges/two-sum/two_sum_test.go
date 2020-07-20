package two_sum_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/two-sum"
)

var _ = Describe("Two Sum", func() {

	Context("two sum general logic", func() {
		When("numbres [2, 7, 11, 15], target: 9", func() {
			It("should return indices: [0, 1]", func() {

				var numbers = []int{2, 7, 11, 15}
				var target = 9

				indices := TwoSum(numbers, target)

				Ω(indices).Should(Equal([]int{0, 1}))
			})
		})

		When("numbres [15, 7, 2, 11], target: 9", func() {
			It("should return indices: [1, 2]", func() {

				var numbers = []int{15, 7, 2, 11}
				var target = 9

				indices := TwoSum(numbers, target)

				Ω(indices).Should(Equal([]int{1, 2}))
			})
		})

		When("numbres [10, 11, 2, 16], target: 12", func() {
			It("should return indices: [0, 3]", func() {

				var numbers = []int{10, 11, 2, 16}
				var target = 12

				indices := TwoSum(numbers, target)

				Ω(indices).Should(Equal([]int{0, 2}))
			})
		})
	})

})
