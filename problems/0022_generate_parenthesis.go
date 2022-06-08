package problems

import "fmt"

func generateParenthesis(n int) []string {
	dp := make([][]string, 9)
	dp[0] = []string{""}
	dp[1] = []string{"()"}

	for i := 2; i <= n; i++ {
		dp[i] = make([]string, 0, 8)
		for j := 0; j < i; j++ {
			for _, p := range dp[j] {
				for _, q := range dp[i-1-j] {
					dp[i] = append(dp[i], fmt.Sprintf("(%s)%s", p, q))
				}
			}
		}
	}

	return dp[n]
}
