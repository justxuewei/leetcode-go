package jzoffer

func reverseWords(s string) string {
	bs := []byte(s)
	input, wordstart := 0, -1
	noword := true

	for i, v := range bs {
		if wordstart == -1 && v != ' ' {
			// mark current position as the start of a word
			noword = false
			wordstart = i
		} else if wordstart != -1 && v == ' ' {
			// reverse the word and insert into the place beginning with i
			reverseWord(bs[wordstart:i])
			for off:=0; off<i-wordstart; off++ {
				bs[input] = bs[wordstart+off]
				input++
			}
			bs[input] = ' '
			input++
			wordstart = -1
		}
	}

	if noword {
		return ""
	}

	if wordstart != -1 {
		reverseWord(bs[wordstart:len(bs)])
		for off:=0; off<len(bs)-wordstart; off++ {
			bs[input] = bs[wordstart+off]
			input++
		}
	} else {
		input--
	}
	bs = bs[:input]
	reverseWord(bs)

	return string(bs)
}

func reverseWord(word []byte) {
	// fmt.Println(string(word))
	for i:=0; i<len(word)/2; i++ {
		word[i], word[len(word)-1-i] = word[len(word)-1-i], word[i]
	}
}

