package main

import (
	"fmt"
	"math"
)

/**
可以将该题转化为用maxCost来获得最长的字符转化
*/

func equalSubstring(s string, t string, maxCost int) int {
	ans := 0
	start, n, cost := 0, len(s), 0
	for end := 0; end < n; end++ {
		cost += int(math.Abs(float64(int(s[end]) - int(t[end]))))
		for cost > maxCost {
			cost -= int(math.Abs(float64(int(s[start]) - int(t[start]))))
			start += 1
		}
		ans = max(ans, end-start+1)
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
	s := "abcd"
	t := "acde"
	maxCost := 0
	fmt.Println(equalSubstring(s, t, maxCost))
}
