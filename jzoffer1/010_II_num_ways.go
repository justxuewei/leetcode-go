package jzoffer1

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	pprev, prev := 1, 1

	for i := 2; i <= n; i++ {
		pprev, prev = prev, (pprev+prev)%(1e9+7)
	}

	return prev
}
