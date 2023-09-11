package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	ans, area := 0, 0
	rows, cols := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		if grid[row][col] == 0 {
			return
		}
		if grid[row][col] == 1 {
			grid[row][col] = 0
			area += 1
		}
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			area = 0
			dfs(i, j)
			ans = max(ans, area)
		}
	}

	return ans
}

func maxAreaOfIsland_bfs(grid [][]int) int {
	ans, area := 0, 0
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				area = bfs(grid, i, j, rows, cols)
			}
			ans = max(ans, area)
		}
	}

	return ans
}

func bfs(grid [][]int, row, col int, rows, cols int) int {
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{row, col})
	ans := 1
	for len(queue) != 0 {
		coor := queue[0]
		queue = queue[1:]
		x, y := coor[0], coor[1]
		if x+1 < rows && grid[x+1][y] == 1 {
			ans += 1
			grid[x+1][y] = 0
			queue = append(queue, [2]int{x + 1, y})
		}
		if x-1 >= 0 && grid[x-1][y] == 1 {
			ans += 1
			grid[x-1][y] = 0
			queue = append(queue, [2]int{x - 1, y})
		}
		if y-1 >= 0 && grid[x][y-1] == 1 {
			ans += 1
			grid[x][y-1] = 0
			queue = append(queue, [2]int{x, y - 1})
		}
		if y+1 < cols && grid[x][y+1] == 1 {
			ans += 1
			grid[x][y+1] = 0
			queue = append(queue, [2]int{x, y + 1})
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
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland_bfs(grid))
}
