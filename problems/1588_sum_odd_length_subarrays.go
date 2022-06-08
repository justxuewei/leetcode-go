package problems

func sumOddLengthSubarrays(arr []int) int {
	var sum, tmp int
	for i := 1; i <= len(arr); i += 2 {
		for j := 0; j < len(arr); j++ {
			tmp = 0
			if j+i > len(arr) {
				continue
			}
			for k := j; k < j+i; k++ {
				tmp += arr[k]
			}
			sum += tmp
		}
	}
	return sum
}
