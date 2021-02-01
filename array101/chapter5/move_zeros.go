package chapter5

func moveZeros(nums []int) {
	j := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	if j < len(nums) {
		for ; j<len(nums); j++ {
			nums[j] = 0
		}
	}
}
