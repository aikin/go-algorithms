package move_zeroes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/array/s1-two-pointer-swap-or-assignment/move-zeroes"
)

var _ = Describe("MoveZeroes", func() {
	Context("move zeroes general logic", func() {
		When("when input nums is : [0, 1, 0, 3, 12]", func() {
			It("should return length: [1, 3, 12, 0, 0]", func() {

				var nums = []int{0, 1, 0, 3, 12}

				MoveZeroes(nums)

				Ω(nums).Should(Equal([]int{1, 3, 12, 0, 0}))
			})
		})

		When("when input nums is : [0, 0, 0, 3, 12, 0, 1]", func() {
			It("should return length: [3, 12, 1, 0, 0, 0, 0]", func() {

				var nums = []int{0, 0, 1, 0, 3, 12, 0, 1}

				MoveZeroes(nums)

				Ω(nums).Should(Equal([]int{1, 3, 12, 1, 0, 0, 0, 0}))
			})
		})

		When("when input nums is : [1, 0, 1]", func() {
			It("should return length: [1, 1, 0]", func() {

				var nums = []int{1, 0, 1}

				MoveZeroes(nums)

				Ω(nums).Should(Equal([]int{1, 1, 0}))
			})
		})
	})
})
