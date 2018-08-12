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

			Ω(list.Length).Should(Equal(1))
		})
	})
})
