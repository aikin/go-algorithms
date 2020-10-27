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
})
