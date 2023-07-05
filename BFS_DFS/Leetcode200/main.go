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
				dfs(grid, i, j, rows, cols)
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

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
