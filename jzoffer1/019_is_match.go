package jzoffer1

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for i:=2; i<=len(p); i=i+2 {
		if !dp[0][i-2] {
			break
		}
		dp[0][i] = p[i-1] == '*'
	}

	for i:=1; i<=len(s); i++ {
		for j:=1; j<=len(p); j++ {
			if p[j-1] == '*' {
				if dp[i][j-2] || dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.') {
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
