package problems

import (
	"fmt"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numStrings := make([]string, len(nums))
	var numString string
	for i, num := range nums {
		numString = strconv.Itoa(num)
		numStrings[i] = numString
	}

	quickSort(numStrings)
	str := strings.Join(numStrings, "")
	var zero int
	for ; zero < len(str)-1; zero++ {
		if str[zero] != '0' {
			break
		}
	}

	return str[zero:]
}

func less(lhs, rhs string) bool {
	lr := fmt.Sprintf("%s%s", lhs, rhs)
	rl := fmt.Sprintf("%s%s", rhs, lhs)

	for i := 0; i < len(lr); i++ {
		if lr[i] != rl[i] {
			return lr[i] < rl[i]
		}
	}

	return false
}

func quickSort(arr []string) {
	if len(arr) <= 1 {
		return
	}

	l, r := 1, len(arr)-1

	for l <= r {
		if less(arr[r], arr[0]) {
			r--
			continue
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
	}

	arr[0], arr[r] = arr[r], arr[0]
	quickSort(arr[:r])
	quickSort(arr[r+1:])
}

func largestNumber1(num []int) string {
	quickSort1(num)

	if num[0] == 0 {
		return "0"
	}
	
	sb := new(strings.Builder)
	for _, n := range num {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}

func less1(lhs, rhs int) bool {
	sx, sy := 10, 10
	for sx <= lhs {
		sx *= 10
	}
	for sy <= rhs {
		sy *= 10
	}
	return rhs+sy*lhs <= lhs+sx*rhs
}

func quickSort1(arr []int) {
	if len(arr) <= 1 {
		return
	}

	l, r := 1, len(arr)-1

	for l <= r {
		if less1(arr[r], arr[0]) {
			r--
			continue
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
	}

	arr[0], arr[r] = arr[r], arr[0]
	quickSort1(arr[:r])
	quickSort1(arr[r+1:])
}
