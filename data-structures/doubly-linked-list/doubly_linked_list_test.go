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

	Context("When append node to linked list", func() {
		It("should append node to linked list", func() {
			list := NewDoublyLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Append(1)

			Ω(list.Len()).Should(Equal(1))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(BeNil())
			Ω(list.Head.Previous).Should(BeNil())
			Ω(list.Head.Value).Should(Equal(1))

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Previous).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(1))
		})
	})
})
