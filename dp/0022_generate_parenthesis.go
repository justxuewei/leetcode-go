package dp

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var dp []string
	for i:=0; i<n; i++ {
		for _, l := range generateParenthesis(i) {
			for _, r := range generateParenthesis(n-1-i) {
				dp = append(dp, fmt.Sprintf("(%s)%s", l, r))
			}
		}
	}

	return dp
}
