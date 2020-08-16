package three_sum_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestThreeSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ThreeSum Suite")
}
