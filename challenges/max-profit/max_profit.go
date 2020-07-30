package max_profit

import (
	"errors"
)

const uintSize = 32 <<(^uint(0) >> 32 & 1)
const (
	maxInt = 1 <<(uintSize - 1) - 1
)

/*
Given prices = [7, 1, 5, 3, 6, 4],

Because maxProfit = prices[4] - prices[1]
return 5

复杂度分析：
 - 时间复杂度：O(n^2)
 - 空间复杂度：O(1)
*/
func MaxProfit(prices []int) (int, error) {
	exceptedMinPrice := 1
	invalidPricesFlag := -1

	maxProfit := 0
	for i := 0; i < len(prices) - 1; i++ {
		for j:= i + 1; j < len(prices); j++ {
			if (prices[i] < exceptedMinPrice || prices[j] < exceptedMinPrice) {
				return invalidPricesFlag, errors.New("prices should >= 1")
			}
			profit := prices[j] - prices[i]
			if (profit > maxProfit) {
				maxProfit = profit
			}
		}
	}
	return maxProfit, nil
}

/*
Given prices = [7, 1, 5, 3, 6, 4],

Because maxProfit = prices[4] - prices[1]
return 5

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(1)
*/
func MaxProfitWithOneLoop(prices []int) (int, error) {
	exceptedMinPrice := 1
	invalidPricesFlag := -1

	minPrice := maxInt
	maxProfit := 0

	for _, curPrice := range prices {
		if (curPrice < exceptedMinPrice) {
			return invalidPricesFlag, errors.New("prices should >= 1")
		}
		if (curPrice < minPrice) {
			minPrice = curPrice
		} else if (curPrice - minPrice > maxProfit) {
			maxProfit = curPrice - minPrice
		}
	}

	return maxProfit, nil
}