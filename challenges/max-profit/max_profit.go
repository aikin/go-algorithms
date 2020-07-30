package max_profit

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
func MaxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices) - 1; i++ {
		for j:= i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if (profit > maxProfit) {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func MaxProfitWithOneLoop(prices []int) int {
	minPrice := maxInt
	maxProfit := 0

	for i := 0; i < len(prices) - 1; i++ {
		if (prices[i] < minPrice) {
			minPrice = prices[i]
		} else if (prices[i] - minPrice > maxProfit) {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}