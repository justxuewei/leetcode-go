package jzoffer

func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	// dp init
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	// dp
	for i := 1; i < n; i++ {
		dp[i] = min3(dp[a]*2, dp[b]*3, dp[c]*5)
		if dp[i] == dp[a]*2 {
			a++
		}
		if dp[i] == dp[b]*3 {
			b++
		}
		if dp[i] == dp[c]*5 {
			c++
		}
	}
	return dp[len(dp)-1]
}

func min3(a, b, c int) (min int) {
	min = a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return
}
