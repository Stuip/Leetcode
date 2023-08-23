package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	maxSide := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSide = max(maxSide, dp[i][0])
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			maxSide = max(maxSide, dp[0][i])
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][col] == '1' {
				dp[row][col] = min(min(dp[row-1][col], dp[row][col-1]), dp[row-1][col-1]) + 1
				maxSide = max(maxSide, dp[row][col])
			}
		}
	}
	return maxSide * maxSide
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}
