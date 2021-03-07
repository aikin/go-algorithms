package radix_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRadixSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RadixSort Suite")
}
