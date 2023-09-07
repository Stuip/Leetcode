package main

import "fmt"

const INF = 0x3f3f3f3f

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for length := 2; length <= n; length++ {
		for left := 1; left+length-1 <= n; left++ {
			right := left + length - 1
			dp[left][right] = INF
			for k := left; k < right; k++ {
				cur := max(dp[left][k-1], dp[k+1][right]) + k
				dp[left][right] = min(dp[left][right], cur)
			}
		}
	}
	return dp[1][n]
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
	n := 10
	fmt.Println(getMoneyAmount(n))
}
