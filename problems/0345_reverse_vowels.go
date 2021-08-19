package problems

func reverseVowels(s string) string {
	i, j := 0, len(s)-1

	strbytes := []byte(s)

	for {
		for i < j && !isVowel(strbytes[i]) {
			i++
		}
		for i < j && !isVowel(strbytes[j]) {
			j--
		}
		if i >= j {
			break
		}
		strbytes[i], strbytes[j] = strbytes[j], strbytes[i]
		i++
		j--
	}

	return string(strbytes)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
