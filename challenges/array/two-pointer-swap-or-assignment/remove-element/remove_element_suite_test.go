package remove_element_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRemoveElement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RemoveElement Suite")
}
