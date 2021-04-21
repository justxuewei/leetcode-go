package jzoffer

var right = [2]int{0,1}
var down = [2]int{1,0}
var left = [2]int{0,-1}
var up = [2]int{-1,0}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 { return nil }
	visited := make([][]int, len(matrix))
	for i := range visited {
		visited[i] = make([]int, len(matrix[i]))
	}
	ret := make([]int, len(matrix)*len(matrix[0]))
	traversal(matrix, visited, 0, 0, 0, right, ret)
	return ret
}

func traversal(matrix, visited [][]int, x, y, i int, dir [2]int, ret []int) {
	for {
		visited[x][y] = 1
		ret[i] = matrix[x][y]
		i++
		if i == len(ret) { return }
		nx, ny := x+dir[0], y+dir[1]
		if nx<0 || nx>=len(matrix) || ny<0 || ny>=len(matrix[0]) || visited[nx][ny]==1 {
			break
		}
		x, y = nx, ny
	}

	switch dir {
	case right:
		dir = down
	case down:
		dir = left
	case left:
		dir = up
	case up:
		dir = right
	}

	traversal(matrix, visited, x+dir[0], y+dir[1], i, dir, ret)
}

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 { return nil }

	ret := make([]int, 0, len(matrix)*len(matrix[0]))
	r, b, l, t := len(matrix[0])-1, len(matrix)-1, 0, 0
	for {
		// from left to right
		for i:=l; i<=r; i++ {
			ret = append(ret, matrix[t][i])
		}
		t++
		if t > b { break }
		// from top to bottom
		for i:=t; i<=b; i++ {
			ret = append(ret, matrix[i][r])
		}
		r--
		if l > r { break }
		// from right to left
		for i:=r; i>=l; i-- {
			ret = append(ret, matrix[b][i])
		}
		b--
		if t > b { break }
		// from bottom to top
		for i:=b; i>=t; i-- {
			ret = append(ret, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}

	return ret
}

