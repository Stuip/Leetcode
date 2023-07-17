package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	// 需要重新设置初始矩阵
	// 一旦遇到障碍的话，那么之后的地方设置为0
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		memo[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		memo[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				memo[i][j] = memo[i-1][j] + memo[i][j-1]
			}
		}
	}
	return memo[m-1][n-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				} else if i > 0 {
					dp[i][j] = dp[i-1][j]
				} else if j > 0 {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
