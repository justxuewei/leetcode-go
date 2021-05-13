package jzoffer

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 { return len(s) }
	begin, end := 0, 1
	substrlen := 0
	for ; end<len(s); end++ {
		if newbegin := isContaining(s, begin, end, s[end]); newbegin != -1 {
			if substrlen < len(s[begin:end]) {
				substrlen = len(s[begin:end])
			}
			begin = newbegin+1
		}
	}
	if substrlen < len(s[begin:end]) {
		substrlen = len(s[begin:end])
	}
	return substrlen
}

func isContaining(s string, begin, end int, c byte) int {
	for i:=begin; i<end; i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// dp
func lengthOfLongestSubstring1(s string) int {
	if len(s) < 2 { return len(s) }

	m := make(map[byte]int)
	var ret, tmp int
	for j := range s {
		i, ok := m[s[j]]
		m[s[j]] = j
		if ok && tmp >= j - i {
			tmp = j - i
		} else {
			tmp++
		}
		if ret < tmp {
			ret = tmp
		}
	}
	return ret
}
