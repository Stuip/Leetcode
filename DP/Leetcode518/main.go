package main

import "fmt"

// Space: O(n)
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		val := coins[i-1]
		for j := 0; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			for k := 1; k*val <= j; k++ {
				dp[i][j] += dp[i-1][j-k*val]
			}
		}
	}
	return dp[n][amount]
}

// Space: O(1)
func change1(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		val := coins[i-1]
		for j := val; j <= amount; j++ {
			dp[j] += dp[j-val]
		}
	}
	return dp[amount]
}

func main() {
	amout := 5
	coins := []int{1, 2, 5}
	fmt.Println(change1(amout, coins))
}
