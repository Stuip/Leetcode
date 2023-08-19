package main

import "fmt"

// Time: O(n), Space: O(n)
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([]int, n)
	dp[0] = prices[0] // dp存储前i股票价格中最低价格
	for i := 1; i < n; i++ {
		dp[i] = min(dp[i-1], prices[i])
	}
	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, prices[i]-dp[i])
	}
	return ans
}

// Space: O(1)
func maxProfit1(prices []int) int {
	n := len(prices)
	ans := 0
	dp := prices[0] // dp存储前i股票价格中最低价格
	for i := 1; i < n; i++ {
		dp = min(dp, prices[i])
		ans = max(ans, prices[i]-dp)
	}

	return ans
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
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
