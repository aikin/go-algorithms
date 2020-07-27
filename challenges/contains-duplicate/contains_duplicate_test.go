package contains_duplicate_test


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/contains-duplicate"
)

var _ = Describe("Contains Duplicate", func() {

	Context("Contains duplicate logic ", func() {
		When("numbres [1, 2, 3, 1]", func() {
			It("should return duplicated: fase", func() {

				var numbers = []int{1, 2, 3, 1}

				// isDuplicated, err := ContainsDuplicate(numbers)

				// isDuplicated, err := ContainsDuplicateAfterSort(numbers)

				isDuplicated, err := ContainsDuplicateWithHashMap(numbers)

				Ω(err).Should(BeNil())
				Ω(isDuplicated).Should(Equal(true))
			})
		})

		When("numbres [1, 2, 3, 4]", func() {
			It("should return duplicated: fase", func() {

				var numbers = []int{1, 2, 3, 4}

				// isDuplicated, err := ContainsDuplicate(numbers)

				// isDuplicated, err := ContainsDuplicateAfterSort(numbers)

				isDuplicated, err := ContainsDuplicateWithHashMap(numbers)


				Ω(err).Should(BeNil())
				Ω(isDuplicated).Should(Equal(false))
			})
		})

		When("numbres [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]", func() {
			It("should return duplicated: false", func() {

				var numbers = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

				// isDuplicated, err := ContainsDuplicate(numbers)

				// isDuplicated, err := ContainsDuplicateAfterSort(numbers)

				isDuplicated, err := ContainsDuplicateWithHashMap(numbers)

				Ω(err).Should(BeNil())
				Ω(isDuplicated).Should(Equal(true))
			})
		})
	})
})