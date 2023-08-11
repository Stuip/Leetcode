package main

import (
	"fmt"
)

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m {
			return 0
		}
		if j < 0 || j >= n {
			return 0
		}
		if used[i][j] || grid[i][j] == 0 {
			return 0
		}
		used[i][j] = true
		res := grid[i][j] + max(max(dfs(i-1, j), dfs(i+1, j)), max(dfs(i, j-1), dfs(i, j+1)))
		used[i][j] = false
		return res
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}
	fmt.Println(getMaximumGold(grid))
}
