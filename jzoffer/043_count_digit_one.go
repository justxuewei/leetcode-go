package jzoffer

func countDigitOne(n int) int {
	var ret int
	digit := 1
	tmp := n
	for tmp > 0 {
		switch tmp % 10 {
		case 0:
			ret += tmp / 10 * digit
		case 1:
			ret += tmp / 10 * digit + n % digit + 1
		default:
			ret += (tmp / 10 + 1) * digit
		}
		digit *= 10
		tmp /= 10
	}
	return ret
}
