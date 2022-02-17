package problems

func balancedStringSplit(s string) (ret int) {
	if len(s) < 2 {
		return 0
	}

	var (
		remaining int
		char      byte
	)

	for i := range s {
		if remaining == 0 {
			char = s[i]
		}

		if s[i] != char {
			remaining--
		} else {
			remaining++
		}

		if remaining == 0 {
			ret++
		}
	}

	return
}
