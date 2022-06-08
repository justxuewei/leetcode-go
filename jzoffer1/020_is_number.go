package jzoffer1

import "strconv"

var fna = []map[string]int{
	{" ": 0, "+/-": 1, "nums": 2, ".": 3},           // 0
	{"nums": 2, ".": 3},                             // 1
	{"nums": 2, ".": 4, "e/E": 5, " ": 8, "end": 9}, // 2
	{"nums": 4},                             // 3
	{"nums": 4, "e/E": 5, " ": 8, "end": 9}, // 4
	{"+/-": 6, "nums": 7},                   // 5
	{"nums": 7},                             // 6
	{"nums": 7, " ": 8, "end": 9},           // 7
	{" ": 8, "end": 9},                      // 8
}

func isNumber(s string) bool {
	var (
		state  int
		cntstr string
	)
	for i := range s {
		switch s[i] {
		case ' ':
			cntstr = " "
		case '+', '-':
			cntstr = "+/-"
		case '.':
			cntstr = "."
		case 'e', 'E':
			cntstr = "e/E"
		default:
			if _, err := strconv.Atoi(string(s[i])); err == nil {
				cntstr = "nums"
			} else {
				return false
			}
		}

		if next, ok := fna[state][cntstr]; ok {
			state = next
		} else {
			return false
		}
	}

	return fna[state]["end"] == 9
}
