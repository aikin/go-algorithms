package quick_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQuickSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "QuickSort Suite")
}
