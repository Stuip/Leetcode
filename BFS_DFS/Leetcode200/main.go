package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ans := 0
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				ans += 1
				bfs(grid, i, j, rows, cols)
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, row, col int, rows, cols int) {
	if row < 0 || col < 0 || row >= rows || col >= cols {
		return
	}
	if grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	dfs(grid, row+1, col, rows, cols)
	dfs(grid, row-1, col, rows, cols)
	dfs(grid, row, col+1, rows, cols)
	dfs(grid, row, col-1, rows, cols)
}

func bfs(grid [][]byte, row, col int, rows, cols int) {
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{row, col})
	for len(queue) != 0 {
		coor := queue[0]
		queue = queue[1:]
		x, y := coor[0], coor[1]
		if x+1 < rows && grid[x+1][y] == '1' {
			grid[x+1][y] = '0'
			queue = append(queue, [2]int{x + 1, y})
		}
		if x-1 >= 0 && grid[x-1][y] == '1' {
			grid[x-1][y] = '0'
			queue = append(queue, [2]int{x - 1, y})
		}
		if y+1 < cols && grid[x][y+1] == '1' {
			grid[x][y+1] = '0'
			queue = append(queue, [2]int{x, y + 1})
		}
		if y-1 >= 0 && grid[x][y-1] == '1' {
			grid[x][y-1] = '0'
			queue = append(queue, [2]int{x, y - 1})
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
