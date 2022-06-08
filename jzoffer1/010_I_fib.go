package jzoffer1

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	var (
		pprev = 0
		prev  = 1
	)

	for i := 2; i <= n; i++ {
		pprev, prev = prev, (pprev+prev)%(1e9+7)
	}

	return prev
}
