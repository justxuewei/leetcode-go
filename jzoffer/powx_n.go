package jzoffer

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		n = -n
		x = 1 / x
	}

	if x == float64(1) {
		return 1
	} else if x == float64(-1) {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	ret := float64(1)
	for i := 0; i < n; i++ {
		if ret == 0 {
			break
		}
		ret *= x
	}
	return ret
}

func myPow1(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}

	ret := float64(1)
	for n > 0 {
		if n&1 == 1 {
			ret *= x
		}
		x *= x
		n >>= 1
	}

	return ret
}
