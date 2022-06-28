package problems

import "math"

func twoEggDrop(n int) int {
	dp := make([][]int, 3)
	for i := 1; i <= 2; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 1; i <= n; i++ {
		dp[1][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[2][i] = min(dp[2][i], max(dp[1][j-1], dp[2][i-j])+1)
		}
	}

	return dp[2][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}