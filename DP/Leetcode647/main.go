package main

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += palindromic(s, i, i) + palindromic(s, i, i+1)
	}
	return res
}

func palindromic(s string, i, j int) int {
	ans := 0
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		ans += 1
		i--
		j++
	}
	return ans
}

func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}
