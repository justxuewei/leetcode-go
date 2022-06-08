package jzoffer

func findRepeatNumber(nums []int) int {
	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			return v
		} else {
			numsMap[v] = true
		}
	}
	return -1
}

func findRepeatNumber1(nums []int) int {
	for i, v := range nums {
		if i == v {
			continue
		}
		if nums[v] == v {
			return v
		}
		nums[i], nums[v] = nums[v], nums[i]
	}
	return -1
}
