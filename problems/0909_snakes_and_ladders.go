package problems

// func snakesAndLadders(board [][]int) int {
// 	n := len(board)
// }

func getCoordination(n, num int) (row int, col int) {
	row = n - num/n - 1
	if row%2 != n%2 {
		if (num / n) * n == 0 {
			col = num - 1
		} else {
			col = (num / n) * n - 1
		}
	} else {
		col = n - (num / n) * n
	}
	return
}

func getNum(n, row, col int) int {
	var (
		revertRow    = n - (row + 1)
		multipler    = 1
		startedValue int
	)
	if revertRow%2 != 0 {
		multipler = -1
		startedValue = (revertRow + 1) * n
	} else {
		startedValue = revertRow*n + 1
	}

	return startedValue + multipler*col
}
