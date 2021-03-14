package remove_duplicates_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/array/s1-two-pointer-swap-or-assignment/remove-duplicates"
)

var _ = Describe("RemoveDuplicates", func() {
	Context("remove duplicates general logic", func() {
		When("when input nums is : [1, 1, 2]", func() {
			It("should return length: 2", func() {

				var nums = []int{1, 1, 2}

				length := RemoveDuplicates(nums)

				Ω(length).Should(Equal(2))
				Ω(nums[:length]).Should(Equal([]int{1, 2}))
			})
		})

		When("when input nums is : [1, 2, 3, 4, 4, 5]", func() {
			It("should return length: 5", func() {

				var nums = []int{1, 2, 3, 4, 4, 5}

				length := RemoveDuplicates(nums)

				Ω(length).Should(Equal(5))
				Ω(nums[:length]).Should(Equal([]int{1, 2, 3, 4, 5}))
			})
		})

		When("when input nums is : [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]", func() {
			It("should return length: 5", func() {

				var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

				length := RemoveDuplicates(nums)

				Ω(length).Should(Equal(5))
				Ω(nums[:length]).Should(Equal([]int{0, 1, 2, 3, 4}))
			})
		})
	})
})
