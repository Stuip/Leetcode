package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	l1, l2 := 0, 0 // 最小值的索引，次小值的索引
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[0][i] = grid[0][i]
		if grid[0][i] < grid[0][l1] {
			l1 = i
			l2 = i
		} else if grid[0][i] < grid[0][l2] {
			l2 = i
		}
	}
	for i := 1; i < n; i++ {
		newL1, newL2 := 0, 0
		for j := 0; j < n; j++ {
			if j == l1 {
				dp[i][j] = grid[i][j] + dp[i][l2]
			} else {
				dp[i][j] = grid[i][j] + dp[i][l1]
			}
			if dp[i][j] < dp[i][newL1] {
				newL1 = j
				newL2 = j
			} else if dp[i][j] < dp[i][newL2] {
				newL2 = j
			}
		}
		l1, l2 = newL1, newL2
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
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(matrix))
}
