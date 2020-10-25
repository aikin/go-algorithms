package container_with_most_water_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/container-with-most-water"
)

var _ = Describe("ContainerWithMostWater", func() {
	Context("Container with most water logic ", func() {
		When("height [1, 8, 6, 2, 5, 4, 8, 3, 7]", func() {
			It("should return max area: 49", func() {

				var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

				maxArea, err := MaxArea(height)

				Ω(err).Should(BeNil())
				Ω(maxArea).Should(Equal(49))
			})
		})
	})
})
