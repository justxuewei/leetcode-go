package problems

func fib(n int) int {
	if n == 0 {
		return 0
	}

	pprev, prev := 1, 1

	for i := 2; i < n; i++ {
		pprev, prev = prev, (pprev+prev)%(1e9+7)
	}

	return prev
}
