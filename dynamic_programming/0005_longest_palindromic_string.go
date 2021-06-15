package dynamic_programming

func longestPalindrome(s string) string {
	ml, mr := 0, 0

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for r:=1; r<len(s); r++ {
		for l:=0; l<r; l++ {
			if s[l]==s[r] && (r-l<=2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l > mr-ml {
					ml, mr = l, r
				}
			}
		}
	}

	return s[ml:mr+1]
}
