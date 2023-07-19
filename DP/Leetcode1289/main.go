package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[0][i] = grid[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt

			for k := 0; k < n; k++ {
				if k != j {
					dp[i][j] = min(dp[i][j], dp[i-1][k]+grid[i][j])
				}
			}
		}
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, dp[n-1][i])
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
		{-73, 61, 43, -48, -36},
		{3, 30, 27, 57, 10},
		{96, -76, 84, 59, -15},
		{5, -49, 76, 31, -7},
		{97, 91, 61, -46, 67},
	}
	fmt.Println(minFallingPathSum(matrix))
}
