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

	It("should be possible to stack objects", func() {
		stack := NewStack()
		firstElement := make(map[string]int)
		secondElement := make(map[string]int)

		firstElement["key1"] = 1
		firstElement["key2"] = 12

		secondElement["key1"] = 1
		secondElement["key2"] = 20

		stack.Push(firstElement)
		stack.Push(secondElement)

		Ω(stack.Len()).Should(Equal(2))
		Ω(stack.IsEmpty()).Should(BeFalse())
	})
})

Context("When pop data from stack", func () {
	It("should update len", func() {
		stack := NewStack()
		stack.Push(1)
		stack.Push(2)



		Ω(stack.Pop()).Should(Equal(2))
		Ω(stack.Len()).Should(Equal(1))

		Ω(stack.Pop()).Should(Equal(1))
		Ω(stack.Len()).Should(Equal(0))
		Ω(stack.IsEmpty()).Should(BeTrue())
	})

	It("should be possible to pop objects", func() {
		stack := NewStack()
		firstElement := make(map[string]int)
		secondElement := make(map[string]int)

		firstElement["key1"] = 1
		firstElement["key2"] = 12

		secondElement["key1"] = 1
		secondElement["key2"] = 20

		stack.Push(firstElement)
		stack.Push(secondElement)

		Ω(stack.Len()).Should(Equal(2))

		f := stack.Pop()
		first := f.(map[string]int)

		Ω(first["key2"]).Should(Equal(20))
		Ω(stack.Len()).Should(Equal(1))

		s := stack.Pop()
		second := s.(map[string]int)

		Ω(second["key2"]).Should(Equal(12))
		Ω(stack.IsEmpty()).Should(BeTrue())
	})
})

Context("When peek data from stack", func() {
	It("should peek top element from stack", func() {
		stack := NewStack()
		stack.Push(3)
		stack.Push(2)

		element := stack.Peek()

		Ω(stack.Len()).Should(Equal(2))
		Ω(element).Should(Equal(2))
	})
})


})
