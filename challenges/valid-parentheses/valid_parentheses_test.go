package valid_parentheses_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/valid-parentheses"
)

var _ = Describe("ValidParentheses", func() {
	Context("valid parenthese logic", func() {
		When("string: ''", func() {
			It("should return valid: true", func() {

				parentheses := ""

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(true))
			})
		})

		When("string: '()('", func() {
			It("should return valid: false", func() {

				parentheses := "()("

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(false))
			})
		})

		When("string: '}{}(){'", func() {
			It("should return valid: false", func() {

				parentheses := "}{}(){"

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(false))
			})
		})

		When("string: (]", func() {
			It("should return valid: false", func() {

				parentheses := "(]"

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(false))
			})
		})

		When("string: ([)]", func() {
			It("should return valid: false", func() {

				parentheses := "([)]"

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(false))
			})
		})

		When("string: ()", func() {
			It("should return valid: true", func() {

				parentheses := "()"

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(true))
			})
		})

		When("string: ()[]{}", func() {
			It("should return valid: true", func() {

				parentheses := "()[]{}"

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(true))
			})
		})

		When("string: {[]}", func() {
			It("should return valid: true", func() {

				parentheses := "{[]}"

				valid, err := IsValid(parentheses)

				Ω(err).Should(BeNil())
				Ω(valid).Should(Equal(true))
			})
		})
	})
})
