package jzoffer

func sumNums(n int) int {
	var f func(ans *int, n int) bool
	f = func(ans *int, n int) bool {
		*ans += n
		return n > 0 && f(ans, n-1)
	}

	var ret int
	f(&ret, n)
	return ret
}
