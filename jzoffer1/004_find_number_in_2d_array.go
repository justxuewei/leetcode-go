package jzoffer1

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var i, j int
	for ; i<len(matrix); i++ {
		if matrix[i][len(matrix[0])-1] >= target {
			break
		}
	}

	for i < len(matrix) {
		
		i++
	}

	return false
}
