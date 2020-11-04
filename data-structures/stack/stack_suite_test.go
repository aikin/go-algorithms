package stack_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}
