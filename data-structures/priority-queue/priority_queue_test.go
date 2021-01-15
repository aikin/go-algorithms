package priority_queue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/data-structures/priority-queue"
)

var _ = Describe("PriorityQueue", func() {
	Context("When create empty hash table", func() {
		It("should create hash table of certain size", func() {
			 hashTable := NewHashTable(1000)

			 Ω(hashTable.Capacity).Should(Equal(1000))
		})
	})
})
