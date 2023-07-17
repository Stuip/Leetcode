package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	// Use Space O(n^2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = triangle[i][j] + min(dp[i-1][j-1], dp[i-1][j])
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt
	for i := 0; i < len(dp[n-1]); i++ {
		ans = min(ans, dp[n-1][i])
	}
	return ans
}

func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	// Space: O(n)
	memo := [2][]int{}
	for i := 0; i < 2; i++ {
		memo[i] = make([]int, n)
	}
	memo[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		cur := i % 2
		prev := 1 - cur
		memo[cur][0] = memo[prev][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			memo[cur][j] = min(memo[prev][j-1], memo[prev][j]) + triangle[i][j]
		}
		memo[cur][i] = memo[prev][i-1] + triangle[i][i]
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, memo[(n-1)%2][i])
	}
	return ans
}

// O(1)
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	for i := 1; i < n; i++ {
		triangle[i][0] += triangle[i-1][0]
		for j := 1; j < i; j++ {
			triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
		}
		triangle[i][i] = triangle[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, triangle[n-1][i])
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
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal2(triangle))
}
