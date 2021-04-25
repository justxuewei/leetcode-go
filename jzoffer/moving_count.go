package jzoffer

var directions = [][]int {
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func movingCount(m int, n int, k int) int {
	// init matrices
	cells := make([][]int, m)
	visited := make([][]bool, m)
	for i:=0; i<m; i++ {
		cells[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	return dfs1(cells, visited, 0, 0, k)
}

func dfs1(cells [][]int, visited [][]bool, x, y, k int) (ncell int) {
	visited[x][y] = true
	// not meet the requirements
	if digitalSum(x) + digitalSum(y) > k {
		return 0
	}
	// dfs
	for _, dir := range directions {
		if isInCell(cells, x+dir[0], y+dir[1]) && !visited[x+dir[0]][y+dir[1]] {
			ncell += dfs1(cells, visited, x+dir[0], y+dir[1], k)
		}
	}
	return ncell+1
}

func isInCell(cells [][]int, x, y int) bool {
	return x >= 0 && x < len(cells) && y >= 0 && y < len(cells[0])
}

func digitalSum(x int) (n int) {
	for x > 0 {
		n += x % 10
		x /= 10
	}
	return
}
