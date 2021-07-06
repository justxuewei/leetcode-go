package dp

func uniquePaths(m, n int) int {
	if m < 2 || n < 2 {
		return 1
	}

	// dp init
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	row, col := 1, 1
	for {
		if row < m {
			for i:=col; i<n; i++ {
				dp[row][i] = dp[row-1][i] + dp[row][i-1]
			}
			row++
		}
		if col < n {
			for i:=row; i<m; i++ {
				dp[i][col] = dp[i-1][col] + dp[i][col-1]
			}
			col++
		}
		if row >= m && col >= n {
			break
		}
	}

	return dp[m-1][n-1]
}
