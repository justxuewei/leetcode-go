package jzoffer1

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1/x
		n = -n
	}

	ret := 1.0
	for n > 0 {
		if n & 1 == 1 {
			ret *= x
		}
		n >>= 1
		x *= x
	}

	return ret
}
