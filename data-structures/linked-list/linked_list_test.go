package list_test

import (
	. "github.com/aikin/go-algorithms/data-structures/linked-list"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	)

var _ = Describe("LinkedList", func() {
	Context("When create empty linked list", func() {
		It("should linked list length equal 0", func() {
			l := NewList()
			Expect(l.Length).To(Equal(0))
		})
	})
})
