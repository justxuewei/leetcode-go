package jzoffer1

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
			return 1
		}
		return 0
	}

	var ret int
	ret = reversePairs(nums[:len(nums)/2]) + reversePairs(nums[len(nums)/2:])

	left, right := deepCopy(nums[:len(nums)/2]), deepCopy(nums[len(nums)/2:])
	var i int
	for ; len(left) > 0 && len(right) > 0; i++ {
		if left[0] <= right[0] {
			nums[i] = left[0]
			left = left[1:]
		} else {
			nums[i] = right[0]
			right = right[1:]
			ret += len(left)
		}
	}

	for ; i < len(nums); i++ {
		if len(left) == 0 {
			nums[i] = right[0]
			right = right[1:]
		} else {
			nums[i] = left[0]
			left = left[1:]
		}
	}

	return ret
}
