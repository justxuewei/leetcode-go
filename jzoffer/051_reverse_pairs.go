package jzoffer

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	tmp := make([]int, len(nums))
	return mergeSort(nums, tmp, 0, len(nums)-1)
}

func mergeSort(nums []int, tmp []int, begin, end int) int {
	if end-begin+1 < 2 {
		return 0
	}
	mid := (begin+end) / 2
	lpairs := mergeSort(nums, tmp, begin, mid)
	rpairs := mergeSort(nums, tmp, mid+1, end)
	// build tmp nums
	for i:=begin; i<=end; i++ {
		tmp[i] = nums[i]
	}
	// find reverse pairs and merge
	var pairs int
	i, j := begin, mid+1
	for k:=begin; k<=end; k++ {
		if i == mid + 1 { // i reaches the end -> do not count pairs
			nums[k] = tmp[j]
			j++
		} else if j == end + 1 || tmp[i] <= tmp[j] { // j reaches the end or tmp[i] less than tmp[j] -> do not count pairs
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
			pairs += mid - i + 1
		}
	}

	return pairs + lpairs + rpairs
}
