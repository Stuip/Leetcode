package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	weight := []int{}
	INF := math.MaxInt32
	// 预处理出所有可以用到的完全平方数
	for t := 1; t*t <= n; t++ {
		weight = append(weight, t*t)
	}
	m := len(weight)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		dp[0][i] = INF
	}

	for i := 1; i <= m; i++ {
		x := weight[i-1]
		for j := 0; j <= n; j++ {
			// 不选择这个数
			dp[i][j] = dp[i-1][j]
			for k := 1; k*x <= j; k++ {
				if dp[i-1][j-k*x] != INF {
					dp[i][j] = min(dp[i][j], dp[i-1][j-k*x]+k)
				}
			}
		}
	}
	return dp[m][n]
}

func numSquares_opt(n int) int {
	INF := math.MaxInt32
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = INF
	}
	for t := 1; t*t <= n; t++ {
		x := t * t
		for j := x; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-x]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares_opt(7929))
}
