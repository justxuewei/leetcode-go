package jzoffer

func minArray(numbers []int) int {
	for i:=1; i<len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return numbers[i]
		}
	}
	return numbers[0]
}
