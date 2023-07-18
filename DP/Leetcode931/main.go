package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for i := 1; i < n; i++ {
		matrix[i][0] += min(matrix[i-1][0], matrix[i-1][1])
		for j := 1; j < n-1; j++ {
			matrix[i][j] += min(min(matrix[i-1][j-1], matrix[i-1][j]), matrix[i-1][j+1])
		}
		matrix[i][n-1] += min(matrix[i-1][n-1], matrix[i-1][n-2])
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, matrix[n-1][i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9}}
	fmt.Println(minFallingPathSum(matrix))
}
