package linked_list_cycle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/list/linked-list-cycle"
	. "github.com/aikin/go-algorithms/data-structures/linked-list"
)

var _ = Describe("LinkedListCycle", func() {
	Context("linked list cycle general logic", func() {
		When("linked list: 3->2->0->-4->2", func() {
			It("should return has cycle: true", func() {
				list := NewLinkedList()

				list.Append(3)
				list.Append(2)
				list.Append(0)
				list.Append(-4)

				foundNode := list.Find(-4)
				foundNode.Next = list.Find(2)

				hasCycle := HasCycle(list.Head)

				Ω(hasCycle).Should(Equal(true))
			})
		})

		When("linked list: 3->2->0->4->2", func() {
			It("should return has cycle: true", func() {
				list := NewLinkedList()

				list.Append(3)
				list.Append(2)
				list.Append(0)
				list.Append(4)

				foundNode := list.Find(4)
				foundNode.Next = list.Find(0)

				hasCycle := HasCycleWithFloyd(list.Head)

				Ω(hasCycle).Should(Equal(true))
			})
		})
	})
})
