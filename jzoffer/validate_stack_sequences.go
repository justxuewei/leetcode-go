package jzoffer

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	stack := make([]int, 0, len(pushed))
	a, b := 0, 0
	for a < len(pushed) {
		if len(stack) != 0 && stack[len(stack)-1] == popped[b] {
			stack = stack[:len(stack)-1]
			b++
		} else {
			stack = append(stack, pushed[a])
			a++
		}
	}
	for len(stack) > 0 && popped[b] == stack[len(stack)-1] {
		stack = stack[:len(stack)-1]
		b++
	}
	return len(stack) == 0
}
