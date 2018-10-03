package doubly_linked_list_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDoublyLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DoublyLinkedList Suite")
}
