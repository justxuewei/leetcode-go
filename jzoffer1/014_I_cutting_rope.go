package jzoffer1

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i:=2; i<=n; i++ {
		for j:=1; j<i; j++ {
			if dp[i] < dp[j] * (i-j) {
				dp[i] = dp[j] * (i-j)
			} else if dp[i] < j * (i-j) {
				dp[i] = j * (i-j)
			}
		}
	}

	return dp[n]
}
