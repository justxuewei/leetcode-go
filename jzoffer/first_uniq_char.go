package jzoffer

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	bmap := make(map[byte]int, len(s))

	for i := range s {
		if _, ok := bmap[s[i]]; ok {
			bmap[s[i]]++
		} else {
			bmap[s[i]] = 1
		}
	}

	for i := range s {
		if times := bmap[s[i]]; times == 1 {
			return s[i]
		}
	}

	return ' '
}

func firstUniqChar1(s string) byte {
	words := [26]int{}

	for i := range s {
		words[s[i]-'a']++
	}

	for i := range s {
		if words[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}
