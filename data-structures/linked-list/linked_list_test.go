package list_test

import (
	. "github.com/aikin/go-algorithms/data-structures/linked-list"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	)

var _ = Describe("LinkedList", func() {
	Context("When create empty linked list", func() {
		It("should linked list length equal 0", func() {
			list := NewLinkedList()
			Ω(list.Len()).Should(BeZero())
		})
	})

	Context("When operate linked list", func() {
		It("should append node to linked list", func() {
			list := NewLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Append(1)

			Ω(list.Len()).Should(Equal(1))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(BeNil())
			Ω(list.Head.Prev).Should(BeNil())
			Ω(list.Head.Value).Should(Equal(1))

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Prev).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(1))
		})
	})
})
