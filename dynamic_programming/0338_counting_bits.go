package dynamic_programming

func countBits(n int) []int {
	ret := make([]int, n+1)
	ret[0] = 0
	for i:=1; i<=n; i++ {
		ret[i] = i & 1 + ret[i >> 1]
	}
	return ret
}
