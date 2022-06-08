package jzoffer

import (
	"math"
)

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	// remove space
	strbytes := []byte(str)
	for i := range strbytes {
		if strbytes[i] != ' ' {
			strbytes = strbytes[i:]
			break
		}
	}
	// sign
	var sign byte
	if strbytes[0] == '+' || strbytes[0] == '-' {
		sign = strbytes[0]
		strbytes = strbytes[1:]
	}
	// bound
	var bound int
	if sign == '-' {
		bound = math.MinInt32 / 10
	} else {
		bound = math.MaxInt32 / 10
	}
	// num
	var ret int
	for i := 0; i < len(strbytes) && strbytes[i] >= '0' && strbytes[i] <= '9'; i++ {
		var tmp int
		if sign == '-' {
			if ret < bound || (ret == bound && int(strbytes[i]-'0') > 8) {
				return math.MinInt32
			}
			tmp = ret*10 - int(strbytes[i]-'0')
		} else {
			if ret > bound || (ret == bound && int(strbytes[i]-'0') > 7) {
				return math.MaxInt32
			}
			tmp = ret*10 + int(strbytes[i]-'0')
		}
		ret = tmp
	}
	return ret
}
