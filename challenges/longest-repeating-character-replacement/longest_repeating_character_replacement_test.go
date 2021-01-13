package longest_repeating_character_replacement_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/longest-repeating-character-replacement"
)

var _ = Describe("LongestRepeatingCharacterReplacement", func() {
	Context("longest repeating character replacement general logic", func() {
		When("s='ABAC', k=2", func() {
			It("should return maxLen: 4", func() {
				s := "ABAC"
				k := 2

				maxLen := CharacterReplacement(s, k)

				Ω(maxLen).Should(Equal(4))
			})
		})

		When("s='AABABBA', k=1", func() {
			It("should return maxLen: 4", func() {
				s := "AABABBA"
				k := 1

				maxLen := CharacterReplacement(s, k)

				Ω(maxLen).Should(Equal(4))
			})
		})

		When("s='AABAAACBDBADD', k=0", func() {
			It("should return maxLen: 4", func() {
				s := "AABAAACBDBADD"
				k := 0

				maxLen := CharacterReplacement(s, k)

				Ω(maxLen).Should(Equal(3))
			})
		})
	})
})
