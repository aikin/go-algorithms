package number_of_islands_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/number-of-islands"
)

var _ = Describe("NumberOfIslands", func() {
	Context("number of is lands general logic", func() {
		When("grid [[1,1,1,1,0], [1,1,0,1,0],[1,1,0,0,0],[0,0,0,0,0]]", func() {
			It("should return numOfIslands : 1", func() {

				var grid = [][]byte {
					{1, 1, 1, 1, 0},
					{1, 1, 0, 1, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0},
				}

				numOfIslands := NumOfIslands(grid)

				Ω(numOfIslands).Should(Equal(1))
			})
		})

		When("grid [[1,1,0,0,0], [1,1,0,0,0],[0,0,1,0,0],[0,0,0,1,1]]", func() {
			It("should return numOfIslands : 3", func() {

				var grid = [][]byte {
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
				}

				numOfIslands := NumOfIslands(grid)

				Ω(numOfIslands).Should(Equal(3))
			})
		})

	})
})

