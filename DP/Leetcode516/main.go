package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for length := 1; length <= n; length++ {
		for left := 0; left+length-1 < n; left++ {
			right := left + length - 1
			if length == 1 {
				dp[left][right] = 1
			} else if length == 2 {
				if s[left] == s[right] {
					dp[left][right] = 2
				} else {
					dp[left][right] = 1
				}
			} else {
				dp[left][right] = max(dp[left+1][right], dp[left][right-1])
				if s[left] == s[right] {
					dp[left][right] = max(dp[left][right], dp[left+1][right-1]+2)
				} else {
					dp[left][right] = max(dp[left][right], dp[left+1][right-1])
				}
			}
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

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}
