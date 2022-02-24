package jzoffer2

func countBits(n int) []int {
	ret := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ret[i] = ret[i/2] + (i % 2)
	}

	return ret
}
