package jzoffer

func isMatch(s string, p string) bool {
	// dp array init
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 2; j <= len(p); j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	// match pattern
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				}
				if dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = true
				}
			} else {
				if dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
