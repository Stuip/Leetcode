package main

import "fmt"

func countGoodSubstrings(s string) int {
	ans := 0
	for i := 0; i <= len(s)-3; i++ {
		m := make(map[byte]int, 3)
		k := i
		for j := 0; j < 3; j++ {
			m[s[k]]++
			k++
		}
		if len(m) == 3 {
			ans += 1
		}
	}
	return ans
}

func main() {
	s := "xyzzaz"
	fmt.Println(countGoodSubstrings(s))
}
