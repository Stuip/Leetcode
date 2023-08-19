package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	Dir := [][]int{
		{0, 1},  // 右移
		{1, 0},  // 下移
		{0, -1}, // 左移
		{-1, 0}, // 上移
	}
	row, col, di := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		newRow, newCol := row+Dir[di][0], col+Dir[di][1]
		if newRow < 0 || newRow >= n || newCol < 0 || newCol >= n || matrix[newRow][newCol] != 0 {
			di = (di + 1) % 4
		}
		row += Dir[di][0]
		col += Dir[di][1]
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
