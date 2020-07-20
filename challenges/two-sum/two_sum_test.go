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

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(indices).Should(Equal([]int{0, 1}))
			})
		})

		When("numbres [15, 7, 2, 11], target: 9", func() {
			It("should return indices: [1, 2]", func() {

				var numbers = []int{15, 7, 2, 11}
				var target = 9

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(indices).Should(Equal([]int{1, 2}))
			})
		})

		When("numbres [10, 11, 2, 16], target: 12", func() {
			It("should return indices: [0, 3]", func() {

				var numbers = []int{10, 11, 2, 16}
				var target = 12

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(indices).Should(Equal([]int{0, 2}))
			})
		})

		When("numbres [2, 4, 6, 0], target: 4", func() {
			It("should return indices: [1, 3]", func() {

				var numbers = []int{2, 4, 6, 0}
				var target = 4

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(indices).Should(Equal([]int{1, 3}))
			})
		})

		When("numbres [2, -7, 11, 15], target: 4", func() {
			It("should return indices: [1, 2]", func() {

				var numbers = []int{2, -7, 11, 15}
				var target = 4

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(indices).Should(Equal([]int{1, 2}))
			})
		})

		When("numbres [2, -7, 11, -15], target: -4", func() {
			It("should return indices: [2, 3]", func() {

				var numbers = []int{2, -7, 11, -15}
				var target = -4

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(indices).Should(Equal([]int{2, 3}))
			})
		})

		When("numbres [2, -7, 11, -15, 6, 5], target: 7", func() {
			It("should return indices: [0, 5]", func() {

				var numbers = []int{2, -7, 11, -15, 6, 5,}
				var target = 7

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(indices).Should(Equal([]int{0, 5}))
			})
		})

		When("numbres [-8, 16, -32, 36], target: -20", func() {
			It("should return indices: [ ]", func() {

				var numbers = []int{-8, 16, -32, 36}
				var target = -20

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(BeNil())
				Ω(len(indices)).Should(BeZero())
			})
		})
	})

	Context("two sum duplicates check logic", func() {
		When("numbres [2, 2, 4, 15], target: 4", func() {
			It("should return indices: null, error: nums duplicates!", func() {

				var numbers = []int{2, 2, 4, 15}
				var target = 4

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(MatchError("nums duplicates"))
				Ω(indices).Should(BeNil())
			})
		})
	})

	When("numbres [2, 3, 15, 15], target: 5", func() {
		It("should return indices: nil, error: nums duplicates", func() {

			var numbers = []int{2, 3, 15, 15}
			var target = 5

			indices, err := TwoSum(numbers, target)

			Ω(err).Should(MatchError("nums duplicates"))
			Ω(indices).Should(BeNil())
		})
	})

	Context("two sum have more one result", func() {
		When("numbres [2, 3, 5, 0], target: 5", func() {
			It("should return indices: nil, error: should exactly one solution", func() {

				var numbers = []int{2, 3, 5, 0}
				var target = 5

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(MatchError("should exactly one solution"))
				Ω(indices).Should(BeNil())
			})
		})

		When("numbres [1, 3, 5, 0, 2], target: 5", func() {
			It("should return indices: nil, error: should exactly one solution", func() {

				var numbers = []int{1, 3, 5, 0, 2}
				var target = 5

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(MatchError("should exactly one solution"))
				Ω(indices).Should(BeNil())
			})
		})
	})

	Context("two sum has overflow size", func() {
		When("numbres [9000000000000000000, 9223372036854775807, -9223372036854775808, 0], target: 9", func() {
			It("should return indices: nil, error: should two sum in [min, max] rang", func() {

				var numbers = []int{9, 9223372036854775807, -9223372036854775808, 0}
				var target = 9

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(MatchError("should two sum in [min, max] rang"))
				Ω(indices).Should(BeNil())

			})
		})

		When("numbres [-9, 9223372036854775807, -9223372036854775808, 0], target: 9", func() {
			It("should return indices: nil, error: should two sum in [min, max] rang", func() {

				var numbers = []int{-9, 9223372036854775807, -9223372036854775808, 0}
				var target = 9

				indices, err := TwoSum(numbers, target)

				Ω(err).Should(MatchError("should two sum in [min, max] rang"))
				Ω(indices).Should(BeNil())

			})
		})
	})
})
