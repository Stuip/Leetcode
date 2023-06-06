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
			if isPalindrome(s, startidx, i) {
				path = append(path, s[startidx:i+1])
				dfs(i + 1)
				path = path[:len(path)-1] // 恢复现场
			}
		}
	}
	dfs(0)
	return ans
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
