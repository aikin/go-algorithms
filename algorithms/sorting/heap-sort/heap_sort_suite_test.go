package heap_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHeapSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HeapSort Suite")
}
