package three_sum_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/three-sum"
)

var _ = Describe("ThreeSum", func() {
	Context("three sum general logic", func() {
		When("numbres [-1, 0, 1, 2, -1, -4]", func() {
			It("should return indices: [[-1, 0, 1], [-1, -1, 2]]", func() {

				var numbers = []int{-1, 0, 1, 2, -1, -4}

				ans := ThreeSum(numbers)

				Ω(ans).Should(Equal([][]int{{-1, -1, 2}, {-1, 0, 1}}))
			})
		})
	})
})
