package remove_nth_from_end_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/remove-nth-from-end"
	. "github.com/aikin/go-algorithms/data-structures/linked-list"
)

var _ = Describe("RemoveNthFromEnd", func() {
	Context("remove nth node from end of linked list general logic", func() {
		When("linked list: 1->2->3->4->5->NULL, n = 2", func() {
			It("should removed linked list: 1->2->3->5->NULL", func() {
				list := NewLinkedList()

				list.Append(1)
				list.Append(2)
				list.Append(3)
				list.Append(4)
				list.Append(5)

				removedHead := RemoveNthFromEnd(list.Head, 2)

				Ω(removedHead).Should(Not(BeNil()))
				Ω(removedHead.Next).Should(Not(BeNil()))
				Ω(removedHead.Value).Should(Equal(1))
				Ω(removedHead.Next.Value).Should(Equal(2))
				Ω(removedHead.Next.Next.Value).Should(Equal(3))
				Ω(removedHead.Next.Next.Next.Value).Should(Equal(5))
			})
		})
	})
})
