package priority_queue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/data-structures/priority-queue"
)

var _ = Describe("PriorityQueue", func() {
	Context("When create min heap priority queue ", func() {
		It("should not be nil", func() {
			priorityQueue := NewMinHeapPriorityQueue()

			Ω(priorityQueue).Should(Not(BeNil()))
		})
	})

	Context("When insert item to min heap priority queue ", func() {
		It("should insert items to the queue and respect priorities", func() {

			priorityQueue := NewMinHeapPriorityQueue()

			priorityQueue.Add(*NewElement(10, 1))
			Ω(priorityQueue.Peek().Value).Should(Equal(10))

			priorityQueue.Add(*NewElement(5, 2))

			Ω(priorityQueue.Peek().Value).Should(Equal(10))

			priorityQueue.Add(*NewElement(100, 0))
			Ω(priorityQueue.Peek().Value).Should(Equal(100))

			Ω(priorityQueue.Len()).Should(Equal(3))
		})

		It("should poll from queue with respect to priorities", func() {
			priorityQueue := NewMinHeapPriorityQueue()

			priorityQueue.Add(*NewElement(10, 1))
			priorityQueue.Add(*NewElement(5, 2))
			priorityQueue.Add(*NewElement(100, 0))
			priorityQueue.Add(*NewElement(200, 0))

			Ω(priorityQueue.Poll().Value).Should(Equal(100))
			Ω(priorityQueue.Poll().Value).Should(Equal(200))
			Ω(priorityQueue.Poll().Value).Should(Equal(10))
			Ω(priorityQueue.Poll().Value).Should(Equal(5))
		})
	})

	Context("When change priority on the min heap priority queue ", func() {
		It("should be possible to change priority of internal nodes", func() {
			priorityQueue := NewMinHeapPriorityQueue()
			priorityQueue.Add(*NewElement(10, 1))
			priorityQueue.Add(*NewElement(5, 2))
			priorityQueue.Add(*NewElement(100, 0))
			priorityQueue.Add(*NewElement(200, 0))

			priorityQueue.ChangePriority(100, 10)
			priorityQueue.ChangePriority(10, 20)

			Ω(priorityQueue.Poll().Value).Should(Equal(200))
			Ω(priorityQueue.Poll().Value).Should(Equal(5))
			Ω(priorityQueue.Poll().Value).Should(Equal(100))
			Ω(priorityQueue.Poll().Value).Should(Equal(10))
		})

		It("should be possible to change priority of head node", func() {
			priorityQueue := NewMinHeapPriorityQueue()
			priorityQueue.Add(*NewElement(10, 1))
			priorityQueue.Add(*NewElement(5, 2))
			priorityQueue.Add(*NewElement(100, 0))
			priorityQueue.Add(*NewElement(200, 0))

			priorityQueue.ChangePriority(200, 10)
			priorityQueue.ChangePriority(10, 20)

			Ω(priorityQueue.Poll().Value).Should(Equal(100))
			Ω(priorityQueue.Poll().Value).Should(Equal(5))
			Ω(priorityQueue.Poll().Value).Should(Equal(200))
			Ω(priorityQueue.Poll().Value).Should(Equal(10))
		})

		It("should be possible to change priority along with node addition", func() {
			priorityQueue := NewMinHeapPriorityQueue()
			priorityQueue.Add(*NewElement(10, 1))
			priorityQueue.Add(*NewElement(5, 2))
			priorityQueue.Add(*NewElement(100, 0))
			priorityQueue.Add(*NewElement(200, 0))

			priorityQueue.ChangePriority(200, 10)
			priorityQueue.ChangePriority(10, 20)

			priorityQueue.Add(*NewElement(15, 15))

			Ω(priorityQueue.Poll().Value).Should(Equal(100))
			Ω(priorityQueue.Poll().Value).Should(Equal(5))
			Ω(priorityQueue.Poll().Value).Should(Equal(200))
			Ω(priorityQueue.Poll().Value).Should(Equal(15))
			Ω(priorityQueue.Poll().Value).Should(Equal(10))
		})

		It("should reutrn error for item is not in priority", func() {
			priorityQueue := NewMinHeapPriorityQueue()
			priorityQueue.Add(*NewElement(10, 1))
			priorityQueue.Add(*NewElement(5, 2))
			priorityQueue.Add(*NewElement(100, 0))
			priorityQueue.Add(*NewElement(200, 0))

			err := priorityQueue.ChangePriority(111, 10)

			Ω(err).Should(MatchError("Element not found"))
			Ω(priorityQueue.Poll().Value).Should(Equal(100))
			Ω(priorityQueue.Poll().Value).Should(Equal(5))
			Ω(priorityQueue.Poll().Value).Should(Equal(200))
			Ω(priorityQueue.Poll().Value).Should(Equal(15))
			Ω(priorityQueue.Poll().Value).Should(Equal(10))
		})
	})
})
