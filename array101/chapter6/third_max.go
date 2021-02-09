package chapter6

func thirdMax(nums []int) int {
	// length of nums less than 3, return the maximum value
	if len(nums) < 3 {
		max := nums[0]
		for i:=1; i<len(nums);i++ {
			if max < nums[i] {
				max = nums[i]
			}
		}
		return max
	}

	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}

	secondMax := 0
	if nums[0] != max {
		secondMax = nums[0]
	} else {
		secondMax = nums[1]
	}
	for _, v := range nums {
		if max > v && secondMax < v {
			secondMax = v
		}
	}

	thirdMax := 0
	for _, v := range nums {
		if v < secondMax {
			thirdMax = v
		}
	}
	for _, v := range nums {
		if secondMax > v && thirdMax < v {
			thirdMax = v
		}
	}

	return thirdMax
}
