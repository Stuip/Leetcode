package main

import "fmt"

// Space: O(1)
func maxProfit(prices []int) int {
	// dp[i][0]: 表示第i天交易完成之后手里没有股票的最大的利润
	// dp[i][1]: 表示第i天交易完成之后手里持有一支股票的最大利润
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, -prices[0] //第一天没有股票,说明没有买卖; 第一天持有股票则买入了,花掉prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// Space: O(1)
func maxProfit1(prices []int) int {
	// dp[i][0]: 表示第i天交易完成之后手里没有股票的最大的利润
	// dp[i][1]: 表示第i天交易完成之后手里持有一支股票的最大利润
	n := len(prices)

	dp1, dp2 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp1 = max(dp1, dp2+prices[i]) //
		dp2 = max(dp2, dp1-prices[i])
	}
	return dp1
}

// Greedy method
func maxProfit2(prices []int) int {
	ans, n := 0, len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(prices))
}
