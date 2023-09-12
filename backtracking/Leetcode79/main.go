package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	length := len(word)
	var dfs func(i, j int, k int) bool
	dfs = func(i, j int, k int) bool {
		if k == length {
			return true
		}
		if i < 0 || i >= m {
			return false
		}
		if j < 0 || j >= n {
			return false
		}
		if used[i][j] || board[i][j] != word[k] {
			return false
		}
		used[i][j] = true
		res := dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
		used[i][j] = false
		return res
	}
	for i := 0; i < m; i++ {
		for j := 0; j <= n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
