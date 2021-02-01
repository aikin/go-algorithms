package bubble_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBubbleSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BubbleSort Suite")
}
