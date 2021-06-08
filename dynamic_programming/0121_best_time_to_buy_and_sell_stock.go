package dynamic_programming

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	buyday := 0
	ret := 0
	for i:=1; i<len(prices); i++ {
		profit := prices[i] - prices[buyday]
		if  profit < 0  {
			buyday = i
		} else if profit > ret {
			ret = profit
		}
	}
	return ret
}
