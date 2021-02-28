package merge_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMergeSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MergeSort Suite")
}
