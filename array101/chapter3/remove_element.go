package chapter3

func removeElement(nums []int, val int) int {
	newI := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[newI] = nums[i]
			newI++
		}
	}
	return newI + 1
}
