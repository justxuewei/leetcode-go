package problems

import "strings"

func solveNQueens(n int) [][]string {
	var ret [][]string
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	solveNQueensBacktrack(board, 0, &ret)

	return ret
}

func isValid(board [][]string, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func solveNQueensBacktrack(board [][]string, row int, ret *[][]string) {
	if row == len(board) {
		solution := make([]string, len(board))
		for i := 0; i < len(board); i++ {
			solution[i] = strings.Join(board[i], "")
		}
		*ret = append(*ret, solution)
		return
	}
	for col := 0; col < len(board[0]); col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		solveNQueensBacktrack(board, row+1, ret)
		board[row][col] = "."
	}
}
