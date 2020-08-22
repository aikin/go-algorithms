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

	Context("When prepend node to linked list", func() {
		It("should prepend node to linked list", func() {
			list := NewLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Prepend(2)

			Ω(list.Len()).Should(Equal(1))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(BeNil())
			Ω(list.Head.Value).Should(Equal(2))

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(2))
		})

		It("should prepend multiple node to linked list", func() {
			list := NewLinkedList()

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Prepend(2)
			list.Append(1)
			list.Prepend(3)

			Ω(list.Len()).Should(Equal(3))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(Not(BeNil()))
			Ω(list.Head.Value).Should(Equal(3))

			Ω(list.Head.Next.Value).Should(Equal(2))
			Ω(list.Head.Next.Next).Should(Not(BeNil()))
			Ω(list.Head.Next.Next.Value).Should(Equal(1))

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(1))
		})
	})

	Context("When remove node from linked list", func() {
		It("should return nil if linked list is empty", func() {
			list := NewLinkedList()

			Ω(list.Remove(5)).Should(BeNil())
		})

		It("should return nil if node not exist in linked list", func() {
			list := NewLinkedList()

			list.Append(1)
			list.Append(3)

			Ω(list.Len()).Should(Equal(2))
			Ω(list.Remove(5)).Should(BeNil())
			Ω(list.Len()).Should(Equal(2))
		})

		It("should return single node by value from linked list", func() {
			list := NewLinkedList()

			list.Append(1)
			list.Append(3)
			list.Append(2)

			removedNode := list.Remove(3)

			Ω(removedNode.Value).Should(Equal(3))
			Ω(list.Len()).Should(Equal(2))
		})

		It("should remove duplicate nodes by value from linked list", func() {
			list := NewLinkedList()

			Ω(list.Remove(5)).Should(BeNil())

			list.Append(1)
			list.Append(2)
			list.Append(3)
			list.Append(2)
			list.Append(3)
			list.Append(4)
			list.Append(3)
			list.Append(4)
			list.Append(5)


			Ω(list.Len()).Should(Equal(9))
			Ω(list.Head.Value).Should(Equal(1))
			Ω(list.Tail.Value).Should(Equal(5))

			removedNode := list.Remove(3)

			Ω(removedNode.Value).Should(Equal(3))
			Ω(list.Len()).Should(Equal(6))
			Ω(list.Head.Value).Should(Equal(1))
			Ω(list.Tail.Value).Should(Equal(5))
			Ω(list.Head.Next.Value).Should(Equal(2))
			Ω(list.Head.Next.Next.Value).Should(Equal(2))
			Ω(list.Head.Next.Next.Next.Value).Should(Equal(4))
			Ω(list.Head.Next.Next.Next.Next.Value).Should(Equal(4))
			Ω(list.Head.Next.Next.Next.Next.Next.Value).Should(Equal(5))
			Ω(list.Head.Next.Next.Next.Next.Next.Next).Should(BeNil())
		})

		It("should remove linked list head", func() {
			list := NewLinkedList()

			headNode := 1
			tailNode := 2

			list.Append(headNode)
			list.Append(tailNode)


			Ω(list.Len()).Should(Equal(2))
			Ω(list.Head.Value).Should(Equal(headNode))
			Ω(list.Tail.Value).Should(Equal(tailNode))

			removedNode := list.Remove(headNode)

			Ω(removedNode.Value).Should(Equal(headNode))
			Ω(list.Len()).Should(Equal(1))
			Ω(list.Head.Value).Should(Equal(tailNode))
			Ω(list.Tail.Value).Should(Equal(tailNode))
		})

		It("should remove linked list tail", func() {
			list := NewLinkedList()

			headNode := 1
			tailNode := 2

			list.Append(headNode)
			list.Append(tailNode)


			Ω(list.Len()).Should(Equal(2))
			Ω(list.Head.Value).Should(Equal(headNode))
			Ω(list.Tail.Value).Should(Equal(tailNode))

			removedNode := list.Remove(tailNode)

			Ω(removedNode.Value).Should(Equal(tailNode))
			Ω(list.Len()).Should(Equal(1))
			Ω(list.Head.Value).Should(Equal(headNode))
			Ω(list.Tail.Value).Should(Equal(headNode))
		})
	})

	Context("When find node from linked list", func() {

		It("should find node by value", func() {
			list := NewLinkedList()

			Ω(list.Find(5)).Should(BeNil())

			list.Append(1)
			list.Append(2)
			list.Append(3)


			foundNode := list.Find(2)

			Ω(foundNode).Should(Not(BeNil()))
			Ω(foundNode.Value).Should(Equal(2))
		})
	})

	Context("When operate node is obejct", func() {
		It("should store objects in the linked list", func() {
			list := NewLinkedList()

			type Person struct {
				FirstName  string
				LastName string
			}

		   tom := Person {"Tom", "Lu"}
		   kin := Person {"Kin", "Lu"}

			Ω(list.Head).Should(BeNil())
			Ω(list.Tail).Should(BeNil())

			list.Append(tom)
			list.Prepend(kin)

			Ω(list.Len()).Should(Equal(2))

			Ω(list.Head).Should(Not(BeNil()))
			Ω(list.Head.Next).Should(Not(BeNil()))
			Ω(list.Head.Value).Should(Equal(kin))

			Ω(list.Head.Next.Value).Should(Equal(tom))
			Ω(list.Head.Next.Next).Should(BeNil())

			Ω(list.Tail).Should(Not(BeNil()))
			Ω(list.Tail.Next).Should(BeNil())
			Ω(list.Tail.Value).Should(Equal(tom))
		})
	})

	Context("When get node from linked list", func() {

		It("should can get node by index", func() {
			list := NewLinkedList()

			Ω(list.Get(5)).Should(BeNil())

			list.Append(1)
			list.Append(2)
			list.Append(3)


			foundNode := list.Get(2)

			Ω(foundNode).Should(Not(BeNil()))
			Ω(foundNode.Value).Should(Equal(2))
		})
	})


})
