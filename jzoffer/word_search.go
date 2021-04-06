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
			if string(board[i][j]) == string(word[0]) {
				if searchWord(board, visited, i, j, word, "") {
					return true
				}
			}
		}
	}
	return false
}

func searchWord(board [][]byte, visited [][]bool, i, j int, word, cntword string) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] { return false }
	wordidx := len(cntword)
	if string(word[wordidx]) != string(board[i][j]) {
		return false
	}
	newcntword := cntword + string(board[i][j])
	if newcntword == word {
		return true
	}
	visited[i][j] = true
	if searchWord(board, visited, i-1, j, word, newcntword) ||
		searchWord(board, visited, i+1, j, word, newcntword) ||
		searchWord(board, visited, i, j-1, word, newcntword) ||
		searchWord(board, visited, i, j+1, word, newcntword) {
		return true
	} else {
		visited[i][j] = false
		return false
	}
}
