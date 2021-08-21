package problems

func trap(height []int) int {
	var ret int
	idx := -1
	var monostack []int

	for i, v := range height {
		if v != 0 {
			idx = i
			break
		}
	}

	if idx == -1 || idx == len(height)-1 {
		return ret
	}

	monostack = append(monostack, idx)
	for i := idx + 1; i < len(height); i++ {
		for len(monostack) > 0 && height[i] >= height[monostack[len(monostack)-1]] {
			// the front is existed
			if len(monostack) != 1 {
				ret += (trapMin(height[monostack[len(monostack)-2]], height[i]) - height[monostack[len(monostack)-1]]) *
					(i - monostack[len(monostack)-2] - 1)
			}
			monostack = monostack[:len(monostack)-1]
		}
		monostack = append(monostack, i)
	}

	return ret
}

func trapMin(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
