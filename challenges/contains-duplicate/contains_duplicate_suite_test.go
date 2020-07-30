package contains_duplicate_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwoSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contains Duplicate Suite")
}
