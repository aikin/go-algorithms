package product_of_array_except_self_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductOfArrayExceptSelf Suite")
}
