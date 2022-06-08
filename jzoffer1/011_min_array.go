package jzoffer1

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		return numbers[0]
	}

	prev := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if prev > numbers[i] {
			return numbers[i]
		}
		prev = numbers[i]
	}

	return numbers[0]
}
