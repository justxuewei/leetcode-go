package jzoffer1

var movingCountDirection = [][]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	return movingCountDfs(m, n, visited, 0, 0, k)
}

func movingCountDfs(m, n int, visited [][]bool, i, j, k int) int {
	if k < sumDigital(i)+sumDigital(j) {
		return 0
	}
	var ret int
	visited[i][j] = true
	for _, dir := range movingCountDirection {
		if i+dir[0] >= 0 && i+dir[0] < m && j+dir[1] >= 0 && j+dir[1] < n && !visited[i+dir[0]][j+dir[1]] {
			ret += movingCountDfs(m, n, visited, i+dir[0], j+dir[1], k)
		}
	}

	return ret + 1
}

func sumDigital(n int) (ret int) {
	for n > 0 {
		ret += n % 10
		n /= 10
	}
	return
}
