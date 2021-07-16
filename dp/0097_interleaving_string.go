package dp

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i:=1; i<=len(s1); i++ {
		if s1[i-1] != s3[i-1] {
			break
		}
		dp[i][0] = true
	}
	for j:=1; j<=len(s2); j++ {
		if s2[j-1] != s3[j-1] {
			break
		}
		dp[0][j] = true
	}

	for i:=1; i<=len(s1); i++ {
		for j:=1; j<=len(s2); j++ {
			fmt.Printf("s1[i-1]: %c, s3[i+j-2]: %c\n", s1[i-1], s3[i+j-2])
			if dp[i-1][j] && s1[i-1] == s3[i+j-2] {
				dp[i][j] = true
				continue
			}
			fmt.Printf("s2[j-1]: %c, s3[i+j-2]: %c\n", s2[j-1], s3[i+j-2])
			if dp[i][j-1] && s2[j-1] == s3[i+j-2] {
				dp[i][j] = true
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
