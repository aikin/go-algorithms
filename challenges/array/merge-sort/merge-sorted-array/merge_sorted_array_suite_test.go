package merge_sorted_array_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMergeSortedArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MergeSortedArray Suite")
}
