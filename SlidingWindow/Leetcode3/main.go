package main

import (
	"fmt"
)

/*
* 每次固定住区间的左边界，然后通过不断右移右边界，直到不符合条件（即该区间里面右重复字符）
 */
func lengthOfLongestSubstring1(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		var seen [128]bool
		j := i
		for j < n && !seen[s[j]] {
			seen[s[j]] = true
			j++
		}
		ans = max(ans, j-i)
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
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring1(s))
}
