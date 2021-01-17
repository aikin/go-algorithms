package heap_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/data-structures/heap"
)

var _ = Describe("Heap", func() {
	Context("When create MaxHeap", func() {
		It("should create an empty max heap", func() {
			 heap := NewMaxHeap()

			 Ω(heap.IsEmpty()).Should(Equal(true))
			 Ω(heap.Len()).Should(Equal(0))
		})

		It("should add items to the heap and heapify it up", func() {
			heap := NewMaxHeap()

			heap.Insert(Int(5))
			heap.Insert(Int(10))
			heap.Insert(Int(3))

			Ω(heap.IsEmpty()).Should(Equal(false))
			Ω(heap.Len()).Should(Equal(3))
			Ω(heap.Extract()).Should(Equal(Int(10)))
	 })
	})
})
