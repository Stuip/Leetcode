package main

import "fmt"

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	var combinations []string
	n := len(digits)
	path := []rune{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(path) == n {
			s := string(path)
			combinations = append(combinations, s)
			return
		}
		for _, c := range phoneMap[string(digits[idx])] {
			path = append(path, c)
			dfs(idx + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return combinations

}

func letterCombinations1(digits string) []string {
	ans := []string{}
	n := len(digits)
	if n == 0 {
		return ans
	}
	path := make([]byte, n)
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		if startIdx == n {
			ans = append(ans, string(path))
			return
		}
		for _, c := range phoneMap[string(digits[startIdx])] {
			path[startIdx] = byte(c)
			dfs(startIdx + 1)
		}
	}
	dfs(0)
	return ans
}

func main() {
	digits := ""
	fmt.Println(letterCombinations1(digits))
}
