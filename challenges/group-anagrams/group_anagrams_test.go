package group_anagrams_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/group-anagrams"
)

var _ = Describe("GroupAnagrams", func() {
	Context("group anagrams general logic", func() {
		When("numbres ['eat', 'tea', 'tan', 'ate', 'nat', 'bat']", func() {
			It("should return groupedAnagrams: [['eat', 'tea', 'ate'], ['tan', 'nat'], ['bat']]", func() {

				var anagrams = [] string{"eat", "tea", "tan", "ate", "nat", "bat"}

				groupedAnagrams := GroupAnagrams(anagrams)

				Ω(groupedAnagrams).Should(Equal([][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}))
			})

			When("numbres ['are', 'nat', 'tan', 'ear', 'code', 'era']", func() {
				It("should return groupedAnagrams: [['ate', 'eat', 'tea'], ['nat', 'tan'], ['bat']]", func() {

					var anagrams = [] string{"are", "nat", "tan", "ear", "code", "era"}

					groupedAnagrams := GroupAnagrams(anagrams)

					Ω(groupedAnagrams).Should(Equal([][]string{{"are", "ear", "era"}, {"nat", "tan"}, {"code"}}))
				})
			})
		})
	})
})
