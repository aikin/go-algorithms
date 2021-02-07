package insertion_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInsertionSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InsertionSort Suite")
}
