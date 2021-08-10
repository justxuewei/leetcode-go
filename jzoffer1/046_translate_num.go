package jzoffer1

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	numstr := fmt.Sprintf("%d", num)

	pprev, prev := 1, 1
	for i := 1; i < len(numstr); i++ {
		if n, _ := strconv.Atoi(numstr[i-1:i+1]); numstr[i-1] != '0' && n < 26 {
			pprev, prev = prev, pprev+prev
		} else {
			pprev = prev
		}
	}

	return prev
}
