package main

import "fmt"

func integerBreak(n int) int {
	// dp[i] => 指的是数字i拆分之后最大的乘积
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i < n+1; i++ {
		for j := 1; j <= i-1; j++ {
			dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(integerBreak(10))
}
