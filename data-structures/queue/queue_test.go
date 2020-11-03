package queue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/data-structures/queue"
)

var _ = Describe("Queue", func() {
	Context("When create queue", func() {

		It("should be empty", func() {
			queue := NewQueue()

			Ω(queue).Should(Not(BeNil()))
			Ω(queue.Len()).Should(BeZero())
			Ω(queue.IsEmpty()).Should(BeTrue())
		})

	})

	Context("When enqueue data to queue", func() {

		It("should update queue len", func() {
			queue := NewQueue()

			queue.Enqueue(1)
			queue.Enqueue(2)

			Ω(queue.Len()).Should(Equal(2))
			Ω(queue.IsEmpty()).Should(BeFalse())
		})

		It("should be possible to enqueue objects", func() {
			queue := NewQueue()

			firstElement := make(map[string]int)
			secondElement := make(map[string]int)

			firstElement["key1"] = 1
			firstElement["key2"] = 12

			secondElement["key1"] = 1
			secondElement["key2"] = 20


			queue.Enqueue(firstElement)
			queue.Enqueue(secondElement)

			Ω(queue.Len()).Should(Equal(2))
			Ω(queue.IsEmpty()).Should(BeFalse())
		})

	})

	Context("When dequeue data from queue", func() {
		It("should update queue len", func() {
			queue := NewQueue()

			queue.Enqueue(1)
			queue.Enqueue(2)

			Ω(queue.Len()).Should(Equal(2))

			Ω(queue.Dequeue()).Should(Equal(1))
			Ω(queue.Len()).Should(Equal(1))

			Ω(queue.Dequeue()).Should(Equal(2))
			Ω(queue.Len()).Should(Equal(0))
			Ω(queue.IsEmpty()).Should(BeTrue())
		})

		It("should be possible to dequeue objects from queue", func() {
			queue := NewQueue()

			firstElement := make(map[string]int)
			secondElement := make(map[string]int)
			firstElement["key1"] = 1
			firstElement["key2"] = 12
			secondElement["key1"] = 1
			secondElement["key2"] = 20


			queue.Enqueue(firstElement)
			queue.Enqueue(secondElement)

			Ω(queue.Len()).Should(Equal(2))

			f := queue.Dequeue()
			first := f.(map[string]int)

			Ω(first["key1"]).Should(Equal(1))

			Ω(queue.Len()).Should(Equal(1))

			s := queue.Dequeue()
		  second := s.(map[string]int)
			Ω(second["key2"]).Should(Equal(20))

			Ω(queue.Dequeue()).Should(BeNil())
			Ω(queue.IsEmpty()).Should(BeTrue())
		})
	})

	Context("When peek data from queue", func() {
		It("should return top elememt of queue with not empty queue", func() {
			queue := NewQueue()

			queue.Enqueue(1)
			queue.Enqueue(2)

			Ω(queue.Peek()).Should(Equal(1))
			Ω(queue.Peek()).Should(Equal(1))
			Ω(queue.IsEmpty()).Should(BeFalse())

		})

		It("should return nil with empty queue", func() {
			queue := NewQueue()

			Ω(queue.Peek()).Should(BeNil())
		})
	})
})
