package valid_anagram_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwoSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Valid Anagram Suite")
}
