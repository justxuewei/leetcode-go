package dynamic_programming

func longestPalindrome(s string) string {
	maxidx := 0
	dp := make([]int, len(s))
	dp[0] = 1

	var j int
	for i:=1; i<len(s); i++ {
		if i - dp[i-1] - 1 >= 0 && s[i] == s[i-dp[i-1]-1] {
			dp[i] = max(dp[i], dp[i-1] + 2)
		}
		j = i - 1
		for  j>=0 && s[i] == s[j] {
			j--
		}
		dp[i] = max(dp[i], i - j)
		if i-2 >= 0 && s[i-2] == s[i] {
			dp[i] = max(dp[i], 3)
		}
		if dp[i] > dp[maxidx] {
			maxidx = i
		}
	}
	return s[maxidx-dp[maxidx]+1:maxidx+1]
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
