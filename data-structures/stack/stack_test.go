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

			Ω(stack.Len()).Should(Equal(0))
			Ω(stack.IsEmpty()).Should(BeTrue())
	})
})


})
