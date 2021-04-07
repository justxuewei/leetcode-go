package jzoffer

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 { return false }

	hlen, vlen := len(board[0]), len(board)
	// init visited
	visited := make([][]bool, vlen)
	for i := range visited {
		visited[i] = make([]bool, hlen)
	}
	for i:=0; i<vlen; i++ {
		for j:=0; j<hlen; j++ {
			if searchWord(board, visited, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte, visited [][]bool, i, j int, word string, wordlen int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] { return false }
	if string(word[wordlen]) != string(board[i][j]) {
		return false
	}
	newwordlen := wordlen+1
	if newwordlen == len(word) {
		return true
	}
	visited[i][j] = true
	if searchWord(board, visited, i-1, j, word, newwordlen) ||
		searchWord(board, visited, i+1, j, word, newwordlen) ||
		searchWord(board, visited, i, j-1, word, newwordlen) ||
		searchWord(board, visited, i, j+1, word, newwordlen) {
		return true
	} else {
		visited[i][j] = false
		return false
	}
}

var dir = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
	[]int{-1, 0},
}

func exist1(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 { return false }

	hlen, vlen := len(board[0]), len(board)
	// init visited
	visited := make([][]bool, vlen)
	for i := range visited {
		visited[i] = make([]bool, hlen)
	}
	for i:=0; i<vlen; i++ {
		for j:=0; j<hlen; j++ {
			if searchWord1(board, visited, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}
func searchWord1(board [][]byte, visited [][]bool, i, j int, word string, wordlen int) bool {
	if wordlen == len(word)-1 {
		return word[wordlen] == board[i][j]
	}

	if word[wordlen] == board[i][j] {
		visited[i][j] = true
		for _, v := range dir {
			newi := i + v[0]
			newj := j + v[1]
			if newi >= 0 && newj >= 0 && newi < len(board) && newj < len(board[0]) &&
				!visited[newi][newj] &&
				searchWord1(board, visited, newi, newj, word, wordlen+1) {
				return true
			}
		}
		visited[i][j] = false
	}
	return false
}
