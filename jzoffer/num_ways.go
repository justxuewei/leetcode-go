package jzoffer

func numWays(n int) int {
	if n <= 1 { return 1 }
	steps := []int{1, 2}
	for i:=3; i<=n; i++ {
		steps[0], steps[1] = steps[1], (steps[0] + steps[1])%(1e9+7)
	}
	return steps[1]
}
