package valid_anagram_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/valid-anagram"
)

var _ = Describe("ValidAnagram", func() {

	When("s: length, t: leng", func() {
		It("should return: false", func() {
				// isAnagram := IsAnagram("length", "leng")
				isAnagram := IsAnagramConcurrency("length", "leng")
				Ω(isAnagram).Should(Equal(false))
		})
	})

	When("s: anagram, t: nagaram", func() {
		It("should return: true", func() {
				// isAnagram := IsAnagram("anagram", "nagaram")
				isAnagram := IsAnagramConcurrency("anagram", "nagaram")
				Ω(isAnagram).Should(Equal(true))
		})
	})

	When("s: rat, t: car", func() {
		It("should return: false", func() {
				// isAnagram := IsAnagram("rat", "car")
				isAnagram := IsAnagramConcurrency("rat", "car")
				Ω(isAnagram).Should(Equal(false))
		})
	})
})
