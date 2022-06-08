package jzoffer

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit, lowestPrice := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if profit <= prices[i]-lowestPrice {
			profit = prices[i] - lowestPrice
		}
		if lowestPrice >= prices[i] {
			lowestPrice = prices[i]
		}
	}
	return profit
}
