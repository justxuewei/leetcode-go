package jzoffer

func reverseLeftWords(s string, n int) string {
	modn := n % len(s)
	if modn < 1 {
		return s
	}

	ret := make([]byte, len(s))
	for i := range ret {
		if i < len(s)-modn {
			ret[i] = s[i+modn]
		} else {
			ret[i] = s[i-(len(s)-modn)]
		}
	}

	return string(ret)
}

func reverseLeftWords1(s string, n int) string {
	return s[n%len(s):] + s[:n%len(s)]
}
