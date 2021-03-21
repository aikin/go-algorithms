package container_with_most_water

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContainerWithMostWater", func() {
	Context("Container with most water logic ", func() {
		When("height [1, 8, 6, 2, 5, 4, 8, 3, 7]", func() {
			It("should return max area: 49", func() {

				var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

				//maxArea, err := MaxAreaByBruteForceWay(height)
				maxArea, err := MaxAreaByTwoPointerWay(height)

				Ω(err).Should(BeNil())
				Ω(maxArea).Should(Equal(49))
			})
		})
	})
})
