package main

import "fmt"

func longestPalindrome(s string) string {
	n := 0
	start := 0
	for i := 0; i < len(s); i++ {
		cur := max(getLen(s, i, i), getLen(s, i, i+1))
		if cur > n {
			n = cur
			start = i - (n-1)/2
		}
	}
	return s[start : start+n]
}

func getLen(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
