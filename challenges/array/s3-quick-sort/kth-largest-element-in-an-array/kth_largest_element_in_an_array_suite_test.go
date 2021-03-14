package kth_largest_element_in_an_array_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKthLargestElementInAnArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KthLargestElementInAnArray Suite")
}
