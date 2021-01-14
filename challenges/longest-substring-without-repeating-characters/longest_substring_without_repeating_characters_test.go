package longest_substring_without_repeating_characters_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/longest-substring-without-repeating-characters"
)

var _ = Describe("LongestSubstringWithoutRepeatingCharacters", func() {
	Context("longest substring without repeating character general logic", func() {
		When("s='abcabcbb'", func() {
			It("should return len: 3", func() {
				s := "abcabcbb"

				len := LengthOfLongestSubstringByBruteForceWay(s)

				Ω(len).Should(Equal(3))
			})
		})

		When("s='bbbbb'", func() {
			It("should return len: 1", func() {
				s := "bbbbb"

				len := LengthOfLongestSubstringByBruteForceWay(s)

				Ω(len).Should(Equal(1))
			})
		})

		When("s='pwwkew'", func() {
			It("should return len: 3", func() {
				s := "pwwkew"

				len := LengthOfLongestSubstringByBruteForceWay(s)

				Ω(len).Should(Equal(3))
			})
		})

		When("s=''", func() {
			It("should return len: 0", func() {
				s := ""

				len := LengthOfLongestSubstringByBruteForceWay(s)

				Ω(len).Should(Equal(0))
			})
		})
	})
})
