package remove_element_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/array/two-pointer-swap-or-assignment/remove-element"
)

var _ = Describe("RemoveElement", func() {
	Context("remove element general logic", func() {
		When("when input nums: [3, 2, 2, 3], and val: 3", func() {
			It("should return length: 2", func() {

				var nums = []int{3, 2, 2, 3}

				length := RemoveElement(nums, 3)

				Ω(length).Should(Equal(2))
			})
		})

		When("when input nums: [0, 1, 2, 2, 3, 0, 4, 2], and val: 2", func() {
			It("should return length: 2", func() {

				var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}

				length := RemoveElement(nums, 2)

				Ω(length).Should(Equal(5))
			})
		})
	})
})
