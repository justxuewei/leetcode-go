package array101

func findNumbers(nums []int) int {
	num := 0
	for _, v := range nums {
		if v < 0 {
			v = -v
		} else if v == 0 {
			num++
		}

		i := 0
		tmp := v
		for tmp > 0 {
			tmp /= 10
			i++
		}
		if i % 2 == 0 { num++ }
	}
	return num
}
