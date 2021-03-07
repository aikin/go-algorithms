package counting_sort_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCountingSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountingSort Suite")
}
