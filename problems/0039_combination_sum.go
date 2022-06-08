package problems

func combinationSum(candidates []int, target int) [][]int {
	var (
		ret  [][]int
		path []int
	)
	combinationSumDfs(candidates, 0, target, path, &ret)

	return ret
}

func combinationSumDfs(candidates []int, begin, remaining int, path []int, ret *[][]int) {
	if remaining < 0 {
		return
	}

	if remaining == 0 {
		*ret = append(*ret, deepCopy(path))
	}

	path = append(path, 0)

	for i := begin; i < len(candidates); i++ {
		path[len(path)-1] = candidates[i]
		combinationSumDfs(candidates, i, remaining-candidates[i], path, ret)
	}
}

func deepCopy(arr []int) []int {
	ret := make([]int, len(arr))
	for i := range arr {
		ret[i] = arr[i]
	}
	return ret
}
