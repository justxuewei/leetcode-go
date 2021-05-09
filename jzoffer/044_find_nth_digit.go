package jzoffer

func findNthDigit(n int) int {
	if n < 10 { return n }

	usedDigits := 1
	mult := 1
	nDigits := 1

	remaining := n
	for {
		usedDigits += nDigits * mult * 9
		if n -usedDigits < 0 {
			break
		}
		remaining = n - usedDigits
		mult *= 10
		nDigits++
	}

	num := mult + remaining /nDigits
	digit := nDigits - remaining %nDigits - 1

	for digit > 0 {
		num /= 10
		digit--
	}

	return num % 10
}
