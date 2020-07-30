package max_profit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/max-profit"
)

var _ = Describe("MaxProfit", func() {
	Context("max profit logic", func() {
		When("prices: [7, 1, 5, 3, 6, 4]", func() {
			It("should return profit: 5", func() {

				var prices = []int{7, 1, 5, 3, 6, 4}

				//  profit, err := MaxProfit(prices)
				 profit, err := MaxProfitWithOneLoop(prices)

				Ω(profit).Should(Equal(5))
				Ω(err).Should(BeNil())
			})
		})

		When("prices: [7, 6, 4, 3, 1]", func() {
			It("should return profit: 0", func() {

				var prices = []int{7, 6, 4, 3, 1}

				//  profit, err := MaxProfit(prices)
				 profit, err := MaxProfitWithOneLoop(prices)


				Ω(profit).Should(Equal(0))
				Ω(err).Should(BeNil())
			})
		})

		When("prices: [7, 1, -5, -3, 6, 4]", func() {
			It("should return profit: 5", func() {

				var prices = []int{7, 1, -5, -3, 6, 4}

				//  profit, err := MaxProfit(prices)
				 profit, err := MaxProfitWithOneLoop(prices)

				Ω(profit).Should(Equal(-1))
				Ω(err).Should(MatchError("prices should >= 1"))
			})
		})
	})
})
