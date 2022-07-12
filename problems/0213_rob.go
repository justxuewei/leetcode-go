package problems

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max213(nums[0], nums[1])
	}

	return max213(dp213(nums[1:]), dp213(nums[:len(nums)-1]))
}

func dp213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max213(nums[0], nums[1])
	}

	pre, cur := nums[0], max213(nums[0], nums[1])

	for i:=2; i<len(nums); i++ {
		pre, cur = cur, max213(pre+nums[i], cur)
	}

	return cur
}

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}
