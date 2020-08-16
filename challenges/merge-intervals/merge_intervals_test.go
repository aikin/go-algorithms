package merge_intervals_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/merge-intervals"
)

var _ = Describe("MergeIntervals", func() {
	Context("merge intervals general logic", func() {
		When("intervals [[1, 3], [2, 6], [8, 10], [15, 18]]", func() {
			It("should return merged intervals : [[1, 6], [8, 10], [15, 18]]", func() {

				var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

				mergedIntervals := Merge(intervals )

				Ω(mergedIntervals).Should(Equal([][]int{{1, 6}, {8, 10}, {15, 18}}))
			})
		})

		When("intervals [[1, 9], [2, 5], [19, 20], [10, 11], [12, 20], [0, 3], [0, 1], [0, 2]]", func() {
			It("should return merged intervals: [[1, 6], [8, 10], [15, 18]]", func() {

				var intervals = [][]int{{1, 9}, {2, 5}, {19, 20}, {10, 11}, {12, 20}, {0, 3}, {0, 1}, {0, 2}}

				mergedIntervals := Merge(intervals )

				Ω(mergedIntervals).Should(Equal([][]int{{0, 9}, {10, 11}, {12, 20}}))
			})
		})
	})
})
