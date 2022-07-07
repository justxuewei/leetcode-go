package problems

// My solution
func maxMatrixSum(matrix [][]int) int64 {
	var (
		negativeNum int
		lessNumber  int
		tmpPositiveNumber int
		ret int
	)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < 0 {
				negativeNum++
				tmpPositiveNumber = -matrix[i][j]
			} else {
				tmpPositiveNumber = matrix[i][j]
			}

			if lessNumber > tmpPositiveNumber {
				lessNumber = tmpPositiveNumber
			}
			ret += tmpPositiveNumber
		}
	}

	if negativeNum%2 != 0 {
		ret -= 2*lessNumber
	}

	return int64(ret)
}
