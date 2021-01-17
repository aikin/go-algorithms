package palindromic_substrings_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPalindromicSubstrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PalindromicSubstrings Suite")
}
