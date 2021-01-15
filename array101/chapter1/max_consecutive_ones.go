package array101

func findMaxConsecutiveOnes(nums []int) int {
	var startIndex int
	maxLength := 0
	last := false

	for i, v := range nums {

		if v == 1 && !last {
			startIndex = i
			last = true
		}

		if v == 0 && last {
			length := i - startIndex
			if maxLength < length {
				maxLength = length
			}
			last = false
		}

	}

	if last {
		length := len(nums) - startIndex
		if maxLength < length {
			maxLength = length
		}
	}

	return maxLength
}
