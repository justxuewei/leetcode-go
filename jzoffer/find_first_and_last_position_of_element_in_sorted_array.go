package jzoffer

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 { return []int{-1, -1} }
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			i, j := mid, mid
			for i > 0 && nums[i-1] == target {
				i--
			}
			for j < len(nums)-1 && nums[j+1] == target {
				j++
			}
			return []int{i, j}
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return []int{-1, -1}
}
