package bytedance

import "math"

func pacificAtlantic(heights [][]int) [][]int {
	nRow, nCol := len(heights), len(heights[0])
	if nRow == 0 || nCol == 0 { return nil }
	pacific := make([][]bool, nRow)
	atlantic := make([][]bool, nRow)
	for i:=0; i<nRow; i++ {
		pacific[i] = make([]bool, nCol)
		atlantic[i] = make([]bool, nCol)
	}
	for i:=0; i<nRow; i++ {
		dfsFlow(heights, pacific, i, 0, math.MinInt32)
		dfsFlow(heights, atlantic, i, nCol-1, math.MinInt32)
	}
	for j:=0; j<nCol; j++ {
		dfsFlow(heights, pacific, 0, j, math.MinInt32)
		dfsFlow(heights, atlantic, nRow-1, j, math.MinInt32)
	}
	var ret [][]int
	for i:=0; i<nRow; i++ {
		for j:=0; j<nCol; j++ {
			if pacific[i][j] && atlantic[i][j] {
				ret = append(ret, []int{i, j})
			}
		}
	}
	return ret
}

func dfsFlow(heights [][]int, visited [][]bool, row, col int, lastHeight int) {
	nRow, nCol := len(heights), len(heights[0])
	if row < 0 || col < 0 || row >= nRow || col >= nCol {
		return
	}
	if visited[row][col] || lastHeight > heights[row][col] {
		return
	}
	visited[row][col] = true
	dfsFlow(heights, visited, row+1, col, heights[row][col])
	dfsFlow(heights, visited, row-1, col, heights[row][col])
	dfsFlow(heights, visited, row, col+1, heights[row][col])
	dfsFlow(heights, visited, row, col-1, heights[row][col])
}

//func pacificAtlantic(heights [][]int) [][]int {
//	pacific := make([][]int, len(heights))
//	atlantic := make([][]int, len(heights))
//	// init
//	for i:=0; i<len(heights); i++ {
//		pacific[i] = make([]int, len(heights))
//		atlantic[i] = make([]int, len(heights))
//	}
//	for i:=0; i<len(heights); i++ {
//		pacific[0][i] = 1
//		pacific[i][0] = 1
//		atlantic[len(heights)-1][i] = 1
//		atlantic[i][len(heights)-1] = 1
//	}
//	// flow to pacific
//	for i:=1; i<len(heights); i++ {
//		for j:=1; j<len(heights); j++ {
//			if dfsFlow(heights, pacific, i, j, heights[i][j]) {
//				pacific[i][j] = 1
//			} else {
//				pacific[i][j] = -1
//			}
//		}
//	}
//	// flow to atlantic
//	for i:=1; i<len(heights); i++ {
//		for j:=1; j<len(heights); j++ {
//			if dfsFlow(heights, atlantic, i, j, heights[i][j]) {
//				atlantic[i][j] = 1
//			} else {
//				atlantic[i][j] = -1
//			}
//		}
//	}
//	var ret [][]int
//	// combine answers
//	for i:=0; i<len(heights); i++ {
//		for j:=0; j<len(heights); j++ {
//			if atlantic[i][j] == 1 && pacific[i][j] == 1 {
//				ret = append(ret, []int{i, j})
//			}
//		}
//	}
//	return ret
//}
//
//func dfsFlow(heights [][]int, cache [][]int, r int, c int, lastHeight int) bool {
//	if r < 0 || c < 0 || r >= len(heights) || c >= len(heights) { return false }
//	currentHeight := heights[r][c]
//	if lastHeight <= currentHeight {
//		if cache[r][c] == 1 {
//			return true
//		}  else if cache[r][c] == -1 {
//			return false
//		} else if cache[r][c] == 0 {
//			// dfs
//			return dfsFlow(heights, cache, r-1, c, currentHeight) ||
//						dfsFlow(heights, cache, r, c-1, currentHeight) ||
//						dfsFlow(heights, cache, r+1, c, currentHeight) ||
//						dfsFlow(heights, cache, r, c+1, currentHeight)
//		}
//		panic("the range of cache element is expected from -1 to 1.")
//	}
//	return false
//}
