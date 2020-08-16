package merge_intervals_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/merge-intervals"
)

var _ = Describe("MergeIntervals", func() {
	Context("merge intervals general logic", func() {
		When("numbres [[1, 3], [2, 6], [8, 10], [15, 18]]", func() {
			It("should return : [[1, 6], [8, 10], [15, 18]]", func() {

				var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

				mergedIntervals := Merge(intervals )

				Ω(mergedIntervals).Should(Equal([][]int{{1, 6}, {8, 10}, {15, 18}}))
			})
		})
	})
})
