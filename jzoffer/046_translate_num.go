package jzoffer

import "strconv"

func translateNum(num int) int {
	numstr := strconv.Itoa(num)
	return translateNumBacktracing(numstr, 0)
}

func translateNumBacktracing(numstr string, idx int) int {
	if idx == len(numstr) {
		return 1
	}

	way1paths, way2paths := 0, 0

	way1paths = translateNumBacktracing(numstr, idx+1)

	if numstr[idx] != '0' && idx+2 <= len(numstr) {
		way2, err := strconv.Atoi(numstr[idx : idx+2])
		if err == nil && way2 < 26 {
			way2paths = translateNumBacktracing(numstr, idx+2)
		}
	}

	return way1paths + way2paths
}

// dp
func translateNum1(num int) int {
	p1, p2 := 1, 1
	x, y := num%10, 0
	num /= 10

	for num > 0 {
		x, y = num%10, x
		num /= 10
		if x != 0 && 10*x+y < 26 {
			p1, p2 = p1+p2, p1
		} else {
			p2 = p1
		}
	}

	return p1
}
