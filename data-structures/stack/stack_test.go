package stack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/data-structures/stack"
)

var _ = Describe("Stack", func() {

Context("When create stack ", func() {
	It("should be empty", func() {
			stack := NewStack()

			Ω(stack).Should(Not(BeNil()))
			Ω(stack.Len()).Should(Equal(0))
			Ω(stack.IsEmpty()).Should(BeTrue())
	})
})

Context("When stack data to stack", func () {
	It("should update len", func() {
		stack := NewStack()

		stack.Push(1)

		Ω(stack.Len()).Should(Equal(1))
		Ω(stack.IsEmpty()).Should(BeFalse())
	})
})


})
