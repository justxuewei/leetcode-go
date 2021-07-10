package dp

import "strconv"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	dp0, dp1 := 1, 1
	var dpnow int
	for i:=1; i<len(s); i++ {
		cur, _ := strconv.Atoi(s[i:i+1])
		prev, _ := strconv.Atoi(s[i-1:i+1])

		if cur == 0 {
			if prev > 0 && prev <= 26 {
				// reset
				dp0, dp1 = 0, dp0
				continue
			} else {
				return 0
			}
		}

		dpnow = 0
		if cur > 0 {
			dpnow += dp1
		}
		if prev > 0 && prev <= 26 {
			dpnow += dp0
		}
		dp0, dp1 = dp1, dpnow
	}

	return dp1
}
