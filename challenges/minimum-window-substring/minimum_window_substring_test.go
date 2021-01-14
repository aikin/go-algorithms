package minimum_window_substring_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/minimum-window-substring"
)

var _ = Describe("MinimumWindowSubstring", func() {
	Context("minimum window substring general logic", func() {
		When("s='ADOBECODEBANC', t=ABC", func() {
			It("should return ans: BANC", func() {
				s := "ADOBECODEBANC"
				t := "ABC"

				ans := MinWindow(s, t)

				Ω(ans).Should(Equal("BANC"))
			})
		})
	})
})
