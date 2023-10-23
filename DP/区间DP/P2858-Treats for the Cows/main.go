package main

import "fmt"

/***
https://www.luogu.com.cn/problem/P2858
*/

func main() {
	var n int
	fmt.Scanln(&n)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&prices[i])
	}
	fmt.Println(treats(prices))
}

func treats(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = prices[i] * n // 先将长度为1的情况赋值为最后一个拿的(n * prices[i])
	}
	for length := 2; length <= n; length++ { // 天数为2
		for left := 0; left+length <= n; left++ {
			right := left + length - 1
			age := n - length + 1
			a := prices[left]*age + dp[left+1][right]
			b := prices[right]*age + dp[left][right-1]
			dp[left][right] = max(a, b)
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
