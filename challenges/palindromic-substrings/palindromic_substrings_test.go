package palindromic_substrings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/palindromic-substrings"
)

var _ = Describe("PalindromicSubstrings", func() {
	Context("palindromic sub strings general logic", func() {
		When("when input string is : 'abc'", func() {
			It("should return count : 3", func() {

				var s = "abc"

				count := CountSubstringsByBruteForceWay(s)

				Ω(count).Should(Equal(3))
			})
		})

		When("when input string is : 'aaa'", func() {
			It("should return count : 6", func() {

				var s = "aaa"

				count := CountSubstringsByBruteForceWay(s)

				Ω(count).Should(Equal(6))
			})
		})
	})
})

