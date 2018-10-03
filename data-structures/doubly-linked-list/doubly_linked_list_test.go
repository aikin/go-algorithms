package doubly_linked_list_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/data-structures/doubly-linked-list"
)

var _ = Describe("DoublyLinkedList", func() {
	Context("When create empty doubly linked list", func() {
		It("should linked list length equal 0", func() {
			list := NewDoublyLinkedList()
			Ω(list.Len()).Should(BeZero())
		})
	})
})
