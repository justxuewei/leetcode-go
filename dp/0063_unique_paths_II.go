package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := range obstacleGrid {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := range obstacleGrid {
		if obstacleGrid[i][0] != 0 {
			break
		}
		dp[i][0] = 1
	}
	for i := range obstacleGrid[0] {
		if obstacleGrid[0][i] != 0 {
			break
		}
		dp[0][i] = 1
	}

	row, col := 1, 1
	for {
		if row < len(obstacleGrid) {
			for i := col; i < len(obstacleGrid[0]); i++ {
				if obstacleGrid[row][i] == 0 {
					dp[row][i] = dp[row-1][i] + dp[row][i-1]
				}
			}
			row++
		}

		if col < len(obstacleGrid[0]) {
			for i := row; i < len(obstacleGrid); i++ {
				if obstacleGrid[i][col] == 0 {
					dp[i][col] = dp[i-1][col] + dp[i][col-1]
				}
			}
			col++
		}

		if row >= len(obstacleGrid) && col >= len(obstacleGrid[0]) {
			break
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
