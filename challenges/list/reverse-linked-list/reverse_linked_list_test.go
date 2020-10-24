package reverse_linked_list_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/list/reverse-linked-list"
	. "github.com/aikin/go-algorithms/data-structures/linked-list"
)

var _ = Describe("ReverseLinkedList", func() {
	Context("reverse linked list general logic", func() {
		When("linked list: 1->2->3->4->5->NULL", func() {
			It("should reversed linked list: 5->4->3->2->1->NULL", func() {
				list := NewLinkedList()

				list.Append(1)
				list.Append(2)
				list.Append(3)
				list.Append(4)
				list.Append(5)

				// reveredHead := ReverseLinkedList(list.Head)
				reveredHead := ReverseLinkedListByRecursive(list.Head)

				Ω(reveredHead).Should(Not(BeNil()))
				Ω(reveredHead.Next).Should(Not(BeNil()))
				Ω(reveredHead.Value).Should(Equal(5))
				Ω(reveredHead.Next.Value).Should(Equal(4))
				Ω(reveredHead.Next.Next.Value).Should(Equal(3))
				Ω(reveredHead.Next.Next.Next.Value).Should(Equal(2))
				Ω(reveredHead.Next.Next.Next.Next.Value).Should(Equal(1))
			})
		})
	})
})
