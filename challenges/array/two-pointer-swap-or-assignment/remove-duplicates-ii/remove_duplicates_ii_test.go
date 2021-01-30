package remove_duplicates_ii_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/array/two-pointer-swap-or-assignment/remove-duplicates-ii"
)

var _ = Describe("RemoveDuplicatesII", func() {
	Context("remove duplicate ii general logic", func() {
		When("when input nums: [1, 1, 1, 2, 2, 3]", func() {
			It("should return length: 5, nums: [1, 1, 2, 2, 3]", func() {

				var nums = []int{1, 1, 2, 2, 2, 3}

				length := RemoveDuplicates(nums)

				Ω(length).Should(Equal(5))
				Ω(nums[:length]).Should(Equal([]int{1, 1, 2, 2, 3}))
			})
		})

		When("when input nums: [0, 0, 1, 1, 1, 1, 2, 3, 3]", func() {
			It("should return length: 5, nums: [0, 0, 1, 1, 1, 1, 2, 3, 3]", func() {

				var nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

				length := RemoveDuplicates(nums)

				Ω(length).Should(Equal(7))
				Ω(nums[:length]).Should(Equal([]int{0, 0, 1, 1, 2, 3, 3}))
			})
		})
	})
})
