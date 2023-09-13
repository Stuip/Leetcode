package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1) // dp[i] 表示考虑前i个字符是否能被wordDict拼凑起来
	dp[0] = true
	set := make(map[string]bool, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		set[wordDict[i]] = true
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i && !dp[i]; j++ {
			sub := s[j:i]
			if _, ok := set[sub]; ok {
				dp[i] = dp[j]
			}
		}
	}
	return dp[n]
}

func main() {
	s := "leetcode"
	wordDict := []string{
		"leet", "code",
	}
	fmt.Println(wordBreak(s, wordDict))
}
