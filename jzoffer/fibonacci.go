package jzoffer

func fib(n int) int {
	modn := n % (1e9 + 7)
	fibArr := []int{0, 1}
	if modn < 2 {
		return fibArr[modn]
	}
	for i := 2; i <= modn; i++ {
		v := (fibArr[0] + fibArr[1]) % (1e9 + 7)
		fibArr[0], fibArr[1] = fibArr[1], v
	}
	return fibArr[1]
}
