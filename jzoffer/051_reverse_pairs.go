package jzoffer

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	lpairs := reversePairs(nums[:len(nums)/2])
	rpairs := reversePairs(nums[len(nums)/2:])
	// find pairs for nums
	var pairs int
	tmpidx := len(nums)/2
	for i:=0; i<len(nums)/2; i++ {
		for ; tmpidx < len(nums); tmpidx++ {
			if nums[i] > nums[tmpidx] {
				pairs += len(nums) - tmpidx
				break
			}
		}
	}
	// merge
	

	return lpairs + rpairs + pairs
}
