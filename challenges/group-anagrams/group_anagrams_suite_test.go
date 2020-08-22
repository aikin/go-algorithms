package group_anagrams_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGroupAnagrams(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GroupAnagrams Suite")
}
