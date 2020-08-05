package maximum_subarray_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaximumSubarray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaximumSubarray Suite")
}
