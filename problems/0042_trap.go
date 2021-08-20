package problems

func trap(height []int) int {
	var idx, ret int
	var monostack []int

	for i, v := range height {
		if v != 0 {
			idx = i
			break
		}
	}

	monostack = append(monostack, idx)
	for i:=idx; i<len(height); i++ {
		if height[i] >= height[monostack[len(monostack)-1]] {

		}
	}

	return ret
}
