package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	Dir := [][]int{
		{0, 1},  // 右移
		{1, 0},  // 下移
		{0, -1}, // 左移
		{-1, 0}, // 上移
	}
	m, n := len(matrix), len(matrix[0])
	ans := []int{}
	// 什么时候结束该螺旋
	row, col, di := 0, 0, 0
	for i := 0; i < m*n; i++ {
		// 什么调整方向
		ans = append(ans, matrix[row][col])
		matrix[row][col] = 200
		newRow, newCol := row+Dir[di][0], col+Dir[di][1]
		if newRow < 0 || newRow >= m || newCol < 0 || newCol >= n || matrix[newRow][newCol] == 200 {
			di = (di + 1) % 4
		}
		row += Dir[di][0]
		col += Dir[di][1]
	}
	return ans
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))
}
