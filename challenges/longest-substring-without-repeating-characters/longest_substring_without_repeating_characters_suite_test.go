package longest_substring_without_repeating_characters_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LongestSubstringWithoutRepeatingCharacters Suite")
}
