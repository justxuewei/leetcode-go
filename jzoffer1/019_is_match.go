package jzoffer1

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return false
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
		dp[0][i] = p[i] == '*'
	}

	for i:=1; i<=len(s); i++ {
		dp[i][0] = true
		for j:=1; j<=len(p); j++ {
			
		}
	}

	return dp[len(p)][len(s)]
}
