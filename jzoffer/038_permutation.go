package jzoffer

func permutation(s string) []string {
	var path []byte
	visited := make([]bool, len(s))
	var ret []string

	permutationBacktracking(s, path, visited, &ret)

	return ret
}

func permutationBacktracking(s string, path []byte, visited []bool, ret *[]string) {
	if len(s) == len(path) {
		*ret = append(*ret, string(path))
		return
	}
	pending := make(map[byte]int, 0)
	for i := range s {
		if visited[i] {
			continue
		}
		pending[s[i]] = i
	}

	for b, i := range pending {
		path = append(path, b)
		visited[i] = true
		permutationBacktracking(s, path, visited, ret)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
