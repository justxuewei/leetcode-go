package jzoffer

import "math"

func maxValue(grid [][]int) int {
	// init dp array
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]

	// dp
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			above, left := math.MinInt32, math.MinInt32
			if i-1 >= 0 {
				left = dp[i-1][j]
			}
			if j-1 >= 0 {
				above = dp[i][j-1]
			}
			dp[i][j] = maxInt(left, above) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func maxInt(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
