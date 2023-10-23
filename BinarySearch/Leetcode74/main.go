package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1
	for {
		if row < rows && col >= 0 {
			if matrix[row][col] == target {
				return true
			} else if matrix[row][col] < target {
				row += 1
			} else {
				col -= 1
			}
		} else {
			return false
		}
	}
}

func main() {
	nums := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(findNumberIn2DArray(nums, 11))
}
