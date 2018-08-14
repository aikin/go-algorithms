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

	Context("When append node to linked list", func() {
		It("should append node to linked list", func() {
			list := NewLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Append(1)

			Ω(list.Len()).Should(Equal(1))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(BeNil())
			Ω(list.Head.Value).Should(Equal(1))

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(1))
		})

		It("should append multiple nodes to linked list", func() {
			list := NewLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Append(1)
			list.Append(2)
			list.Append(3)

			Ω(list.Len()).Should(Equal(3))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(Not(BeNil()))
			Ω(list.Head.Value).Should(Equal(1))

			Ω(list.Head.Next.Value).Should(Equal(2))
			Ω(list.Head.Next.Next).Should(Not(BeNil()))
			Ω(list.Head.Next.Next.Value).Should(Equal(3))

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(3))
		})
	})
})
