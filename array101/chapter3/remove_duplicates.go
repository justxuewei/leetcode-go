package chapter3

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	newI := 1
	lastElement := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != lastElement {
			nums[newI] = nums[i]
			lastElement = nums[i]
			newI++
		}
	}
	return newI
}
