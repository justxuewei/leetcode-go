package jzoffer

var fsa = []map[byte]int{
	{' ': 0, 's': 1, 'd': 2, '.': 4},
	{'d': 2, '.': 4},
	{'d': 2, '.': 3, 'e': 5, ' ': 8},
	{'d': 3, 'e': 5, ' ': 8},
	{'d': 3},
	{'s': 6, 'd': 7},
	{'d': 7},
	{'d': 7, ' ': 8},
	{' ': 8},
}

func isNumber(s string) bool {
	state := 0
	for _, b := range s {
		var statecode byte
		if b >= '0' && b <= '9' {
			statecode = 'd'
		} else if b == '+' || b == '-' {
			statecode = 's'
		} else if b == 'e' || b == 'E' {
			statecode = 'e'
		} else {
			statecode = byte(b)
		}

		nextstate, ok := fsa[state][statecode]
		if !ok {
			return false
		}
		state = nextstate
	}
	return state == 2 || state == 3 || state == 7 || state == 8
}
