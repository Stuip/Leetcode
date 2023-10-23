package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n, ans := len(s), 0
	record := make(map[byte]int, n)
	left := 0
	for right := 0; right < n; right++ {
		if idx, ok := record[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		record[s[right]] = right
		ans = max(ans, right-left+1)
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
	s := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
