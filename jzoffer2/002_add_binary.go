package jzoffer2

func addBinary(a, b string) string {
	x, y := atoiInBinary(a), atoiInBinary(b)

	temp1, temp2 := 0, 0

	for y > 0 {
		temp1 = x ^ y
		temp2 = (x & y) << 1
		x, y = temp1, temp2
	}

	return itoaInBinary(x)
}

func atoiInBinary(a string) (ret int) {
	for i, char := range a {
		if char == '1' {
			ret |= 1 << (len(a) - i - 1)
		}
	}
	return
}

func itoaInBinary(a int) string {
	if a == 0 {
		return "0"
	}
	ret := make([]byte, getShift(a))
	for i := 0; a>>i != 0; i++ {
		if a>>i&1 == 1 {
			ret[len(ret)-i-1] = '1'
		} else {
			ret[len(ret)-i-1] = '0'
		}
	}
	return string(ret)
}

func getShift(a int) (shift int) {
	for a>>shift != 0 {
		shift++
	}
	return
}
