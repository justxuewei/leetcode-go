package jzoffer1

func findRepeatNumber(nums []int) int {
	set := make(map[int]bool)

	for _, v := range nums {
		if _, ok := set[v]; ok {
			return v
		}
		set[v] = true
	}

	return 0
}
