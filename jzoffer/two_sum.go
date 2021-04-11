package jzoffer

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 { return nil }
	numsmap := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := numsmap[v]; ok {
			continue
		}
		numsmap[v] = true
	}

	for _, v := range nums {
		av := target - v
		if _, ok := numsmap[av]; ok {
			return []int{v, av}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	if len(nums) < 2 { return nil }
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] + nums[right] > target {
			right--
		} else if nums[left] + nums[right] < target {
			left++
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}
