package jzoffer1

var existDirection = [][]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range board {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := range board {
		for j := range board[0] {
			if word[0] == board[i][j] {
				if existDfs(board, visited, i, j ,0, word) {
					return true
				}
			}
		}
	}

	return false
}

func existDfs(board [][]byte, visited [][]bool, i, j, iword int, word string) bool {
	if iword == len(word)-1 {
		return true
	}
	visited[i][j] = true

	var (
		nexti, nextj int
		ret bool
	)
	for _, dir := range existDirection {
		nexti = dir[0] + i
		nextj = dir[1] + j
		if nexti >= 0 && nexti < len(board) && nextj >= 0 && nextj < len(board[0]) &&
			!visited[nexti][nextj] && board[nexti][nextj] == word[iword+1] {
			ret = ret || existDfs(board, visited, nexti, nextj, iword+1, word)
		}
		if ret {
			return true
		}
	}

	visited[i][j] = false
	return false
}
