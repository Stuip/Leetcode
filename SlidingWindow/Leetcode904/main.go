package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {
	start, ans := 0, 0
	n := len(fruits)
	m := make(map[int]int, n)
	for end := 0; end < n; end++ {
		m[fruits[end]]++
		for len(m) > 2 {
			m[fruits[start]]--
			if m[fruits[start]] == 0 {
				delete(m, fruits[start])
			}
			start += 1
		}
		ans = max(ans, end-start)
	}
	return ans + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	friuts := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(friuts))
}
