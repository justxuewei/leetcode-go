package problems

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	var i, maxstep, next int
	for {
		for j := 0; j <= nums[i] && i+j < len(nums); j++ {
			if i+j+nums[i+j] > maxstep {
				maxstep = i + j + nums[i+j]
				next = i + j
				if maxstep >= len(nums)-1 {
					return true
				}
			}
		}
		if next == i {
			return false
		}
		i = next
	}
}
