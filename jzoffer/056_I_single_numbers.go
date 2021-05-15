package jzoffer

func singleNumbers(nums []int) []int {
	var xor int
	for _, v := range nums {
		xor ^= v
	}

	mask := 0
	for xor & 1 != 1 {
		mask++
		xor >>= 1
	}

	ret := make([]int, 2)
	for _, v := range nums {
		if v & (1 << mask) == 0 {
			ret[0] ^= v
		} else {
			ret[1] ^= v
		}
	}

	return ret
}
