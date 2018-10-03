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

	Context("When append node to doubly linked list", func() {
		It("should append node to doubly linked list", func() {
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

		It("should append multiple nodes to doubly linked list", func() {
			list := NewDoublyLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			firstNodeValue := 1
			secondNodeValue := 2
			thirdNodeValue := 3

			list.Append(firstNodeValue)
			list.Append(secondNodeValue)
			list.Append(thirdNodeValue)

			Ω(list.Len()).Should(Equal(3))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(Not(BeNil()))
			Ω(list.Head.Previous).Should(BeNil())
			Ω(list.Head.Value).Should(Equal(firstNodeValue))

			Ω(list.Head.Next.Value).Should(Equal(secondNodeValue))
			Ω(list.Head.Next.Next).Should(Not(BeNil()))
			Ω(list.Head.Next.Next.Value).Should(Equal(thirdNodeValue))
			Ω(list.Head.Next.Next.Next).Should(BeNil())

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Previous).Should(Not(BeNil()))
			Ω(list.Tail.Value).Should(Equal(thirdNodeValue))

			Ω(list.Tail.Previous.Value).Should(Equal(secondNodeValue))
			Ω(list.Tail.Previous.Previous).Should(Not(BeNil()))	
			Ω(list.Tail.Previous.Previous.Value).Should(Equal(firstNodeValue))
			Ω(list.Tail.Previous.Previous.Previous).Should(BeNil())
		})
	})

	Context("When prepend node to doubly linked list", func() {
		It("should prepend node to doubly linked list", func() {
			list := NewDoublyLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Prepend(2)

			Ω(list.Len()).Should(Equal(1))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(BeNil())
			Ω(list.Head.Previous).Should(BeNil())
			Ω(list.Head.Value).Should(Equal(2))

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Previous).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(2))
		})

		It("should prepend multiple node to doubly linked list", func() {
			list := NewDoublyLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			firstNodeValue := 3
			secondNodeValue := 2
			thirdNodeValue := 1


			list.Prepend(secondNodeValue)
			list.Append(thirdNodeValue)
			list.Prepend(firstNodeValue)

			Ω(list.Len()).Should(Equal(3))


			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(Not(BeNil()))
			Ω(list.Head.Previous).Should(BeNil())
			Ω(list.Head.Value).Should(Equal(firstNodeValue))

			Ω(list.Head.Next.Value).Should(Equal(secondNodeValue))
			Ω(list.Head.Next.Next).Should(Not(BeNil()))
			Ω(list.Head.Next.Next.Value).Should(Equal(thirdNodeValue))
			Ω(list.Head.Next.Next.Next).Should(BeNil())

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Previous).Should(Not(BeNil()))
			Ω(list.Tail.Value).Should(Equal(thirdNodeValue))

			Ω(list.Tail.Previous.Value).Should(Equal(secondNodeValue))
			Ω(list.Tail.Previous.Previous).Should(Not(BeNil()))	
			Ω(list.Tail.Previous.Previous.Value).Should(Equal(firstNodeValue))
			Ω(list.Tail.Previous.Previous.Previous).Should(BeNil())
		})
	})

})
