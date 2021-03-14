package sort_colors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSortColors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SortColors Suite")
}
