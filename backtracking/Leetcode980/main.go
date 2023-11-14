package main

import "fmt"

func uniquePathsIII(grid [][]int) int {
	var cnt, sx, sy int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				cnt += 1
			} else if grid[i][j] == 1 {
				sx, sy = i, j
			}
		}
	}

	var dfs func(x, y, cout int) int
	dfs = func(x, y, count int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] < 0 {
			return 0
		}
		if grid[x][y] == 2 {
			if count == 0 {
				return 1
			}
			return 0
		}
		grid[x][y] = -1
		defer func() { grid[x][y] = 0 }()
		return dfs(x-1, y, count-1) + dfs(x+1, y, count-1) + dfs(x, y-1, count-1) + dfs(x, y+1, count-1)
	}
	return dfs(sx, sy, cnt+1)
}

func main() {
	grid := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}
	fmt.Println(uniquePathsIII(grid))
}
