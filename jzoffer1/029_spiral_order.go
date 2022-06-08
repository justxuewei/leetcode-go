package jzoffer1

const (
	top = iota
	right
	bottom
	left
)

var spiralOrderDirections = [][]int{
	{0, 1},  // go right
	{1, 0},  // go down
	{0, -1}, // go left
	{-1, 0}, // go up
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	ret := make([]int, 0, len(matrix)*len(matrix[0]))
	var i, j int
	border := []int{-1, len(matrix[0]), len(matrix), -1}
	diridx := top

	ret = append(ret, matrix[i][j])

	for border[bottom]-border[top] > 1 || border[right]-border[left] > 1 {
		for i+spiralOrderDirections[diridx][0] > border[top] &&
			i+spiralOrderDirections[diridx][0] < border[bottom] &&
			j+spiralOrderDirections[diridx][1] > border[left] &&
			j+spiralOrderDirections[diridx][1] < border[right] {
			i, j = i+spiralOrderDirections[diridx][0], j+spiralOrderDirections[diridx][1]
			ret = append(ret, matrix[i][j])
		}
		if diridx == top || diridx == left {
			border[diridx]++
		} else {
			border[diridx]--
		}
		diridx = (diridx + 1) % 4
	}

	return ret
}
