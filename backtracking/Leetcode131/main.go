package main

import "fmt"

func partition(s string) [][]string {
	ans := [][]string{}
	path := []string{}
	var dfs func(startidx int)
	dfs = func(startidx int) {
		if startidx >= len(s) {
			copyPath := make([]string, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		for i := startidx; i < len(s); i++ {
			path = append(path, string(s[startidx]))
			if !isPalindrome(path) {
				continue
			}
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

func isPalindrome(s []string) bool {
	if len(s) <= 1 {
		return true
	}
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
