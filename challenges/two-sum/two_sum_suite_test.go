package two_sum_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwoSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Two Sum Suite")
}
