package jzoffer1

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var (
		right = len(matrix[0])-1
	)

	for i:=0 ; i<len(matrix); i++ {
		if matrix[i][right] < target {
			continue
		} else if matrix[i][right] == target {
			return true
		} else {
			for right >= 0 && matrix[i][right] > target {
				right--
			}

			if right < 0 {
				return false
			}
			if matrix[i][right] == target {
				return true
			}
		}
	}

	return false
}
