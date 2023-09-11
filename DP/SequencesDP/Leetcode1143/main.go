package main

import "fmt"

//LCS
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] 表示考虑s1的前i个字符,考虑s2字符的前j个字符,形成最长公共子序列长度
	// 如果s1[i] == s2[j]: dp[i][j] = dp[i-1][j-1] + 1
	// 如果s1[i] != s2[j]: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	//   dp[i-1][j]: 必然不使用s1[i]但可能使用s2[j]时
	//   dp[i][j-1]: 必然不使用s2[j]但可能使用s1[i]时
	n, m := len(text1), len(text2)
	text1, text2 = " "+text1, " "+text2
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[n][m] - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text1, text2 := "abcde", "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
