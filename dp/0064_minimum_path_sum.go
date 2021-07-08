package dp

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i:=1; i<n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i:=1; i<m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	row, col := 1, 1

	for {
		if row < m {
			for i:=col; i<n; i++ {
				dp[row][i] = min(dp[row-1][i], dp[row][i-1]) + grid[row][i]
			}
			row++
		}

		if col < n {
			for i:=row; i<m; i++ {
				dp[i][col] = min(dp[i-1][col], dp[i][col-1]) + grid[i][col]
			}
			col++
		}

		if row >= m && col >= n {
			break
		}
	}

	return dp[m-1][n-1]
}
