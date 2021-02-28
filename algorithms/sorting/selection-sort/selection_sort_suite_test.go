package selection_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSelectionSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SelectionSort Suite")
}
