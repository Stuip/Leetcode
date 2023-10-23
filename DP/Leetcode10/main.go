package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s)+1, len(p)+1
	// dp[i][j] 代表 s[:i] 与 p[:j] 是否可以匹配。
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true // 代表两个空字符串可以匹配
	// 初始化首行
	for j := 2; j < n; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[i-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				}
			}
		}
	}
	return false
}

func main() {
	s := "aaa"
	p := "ab*.*"
	fmt.Println(isMatch(s, p))
}
