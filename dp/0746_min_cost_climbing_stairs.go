package dp

func minCostClimbingStairs(cost []int) int {
	if len(cost) > 2 {
		for i := 2; i < len(cost); i++ {
			cost[i] = cost[i] + min(cost[i-2], cost[i-1])
		}
	}

	return min(cost[len(cost)-2], cost[len(cost)-1])
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}
