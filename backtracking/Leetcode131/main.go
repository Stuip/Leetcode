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
				path = path[:len(path)-1]
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

func partitionOther(s string) [][]string {
	path := []string{}
	n := len(s)
	ans := [][]string{}
	var backtrack func(i, startIdx int)
	backtrack = func(i, startIdx int) {
		if i == n {
			copyPath := make([]string, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		// 不在i, i+1之前分割
		if i < n-1 {
			backtrack(i+1, startIdx)
		}
		if isPalindrome(s, startIdx, i) {
			path = append(path, s[startIdx:i+1])
			backtrack(i+1, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return ans
}

func main() {
	fmt.Println(partitionOther("aab"))
}
