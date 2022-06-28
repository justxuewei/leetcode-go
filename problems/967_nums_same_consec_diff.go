package problems

func numsSameConsecDiff(n, k int) []int {
	ret := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tmp int
	power := 1
	for i := 1; i < n; i++ {
		var current []int
		for _, retv := range ret {
			tmp = retv / power
			for j := 0; j < 10; j++ {
				if tmp-j == k || j-tmp == k {
					current = append(current, retv+j*power*10)
				}
			}
		}
		ret = current
		power *= 10
	}

	newRet := []int{}
	for _, retv := range ret {
		if retv/power != 0 {
			newRet = append(newRet, retv)
		}
	}

	return newRet
}
