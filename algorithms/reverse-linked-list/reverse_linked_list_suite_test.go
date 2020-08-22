package reverse_linked_list_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReverseLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReverseLinkedList Suite")
}
