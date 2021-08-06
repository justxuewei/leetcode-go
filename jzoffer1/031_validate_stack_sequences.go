package jzoffer1

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	mockStack := make([]int, 0, len(pushed))

	for len(popped) != 0 {
		if len(mockStack) == 0 || mockStack[len(mockStack)-1] != popped[0] {
			if len(pushed) == 0 {
				return false
			}
			mockStack = append(mockStack, pushed[0])
			pushed = pushed[1:]
		} else {
			mockStack = mockStack[0:len(mockStack)-1]
			popped = popped[1:]
		}
	}

	return true
}
