package jzoffer2

import "strings"

func addBinary(a, b string) string {
	a, b = stringCompletion(a, b)

	sb := strings.Builder{}
	var carry, numa, numb int
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' {
			numa = 0
		} else {
			numa = 1
		}
		if b[i] == '0' {
			numb = 0
		} else {
			numb = 1
		}
		if (numa+numb+carry)%2 == 0 {
			sb.WriteByte('0')
		} else {
			sb.WriteByte('1')
		}
		carry = (numa + numb + carry) / 2
	}

	if carry == 1 {
		sb.WriteByte('1')
	}

	return reverseString(sb.String())
}

func stringCompletion(a, b string) (string, string) {
	if len(a) != len(b) {
		if len(a) < len(b) {
			a, b = b, a
		}
		newb := make([]byte, len(a))
		for i := 0; i < len(a)-len(b); i++ {
			newb[i] = '0'
		}
		for i := 0; i < len(b); i++ {
			newb[len(a)-len(b)+i] = b[i]
		}
		b = string(newb)
	}
	return a, b
}

func reverseString(str string) string {
	ret := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		ret[i] = str[len(str)-1-i]
	}
	return string(ret)
}
