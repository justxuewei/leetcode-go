package bytedance

func subsets(nums []int) [][]int {
	var ret [][]int
	subsetsRecursively(nums, []int{}, &ret)
	return ret
}

func subsetsRecursively(subNums []int, subset []int, ret *[][]int) {
	if len(subNums) == 0 {
		*ret = append(*ret, subset)
		return
	}
	// empty
	subsetsRecursively([]int{}, subset, ret)
	for i:=0; i<len(subNums); i++ {
		var newSubNums []int
		if i + 1 < len(subNums) {
			newSubNums = subNums[i+1:]
		} else {
			newSubNums = []int{}
		}

		newSubSet := append(deepCopy(subset), subNums[i])
		subsetsRecursively(newSubNums, newSubSet, ret)
	}
}

func deepCopy(origin []int) []int {
	ret := make([]int, len(origin))
	for i, v := range origin {
		ret[i] = v
	}
	return ret
}
