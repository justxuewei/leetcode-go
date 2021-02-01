package chapter5

func sortArrayByParity(A []int) []int {

	j := len(A) - 1

	for i:=0; i<j; i++ {
		if A[i] % 2 == 0 { continue }

		for i < j {
			if A[j] % 2 == 0 {
				A[i], A[j] = A[j], A[i]
				break
			}
			j--
		}
	}

	return A
}
