package jzoffer2

import "math"

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	if a == 0 || b == 1 {
		return a
	}
	if b == -1 {
		return -a
	}

	ret := 0
	retMultiplier := 1
	if a < 0 {
		retMultiplier *= -1
		a = -a
	}
	if b < 0 {
		retMultiplier *= -1
		b = -b
	}

	if a == b {
		return retMultiplier
	}

	shift := getMaxShift(a, b)

	for a >= b {
		for a < b<<shift {
			shift--
		}
		a -= b << shift
		ret += 1 << shift
	}

	return retMultiplier * ret
}

func getMaxShift(a, b int) (shift int) {
	for ; a > b<<shift && math.MaxInt32 >= b<<shift; shift++ {
	}
	return
}
