package main

import (
	"fmt"
	"unicode"
)

func letterCasePermutation(s string) []string {
	ans, path := []string{}, []byte{}
	n := len(s)
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if len(path) == n {
			str := string(path)
			ans = append(ans, str)
			return
		}
		if unicode.IsDigit(rune(s[startIdx])) {
			path = append(path, s[startIdx])
			backtrack(startIdx + 1)
			path = path[:len(path)-1]
		} else {
			// 将字母字符 => 转化为小写字母和大写字母拼接
			upper := unicode.ToUpper(rune(s[startIdx]))
			lower := unicode.ToLower(rune(s[startIdx]))
			r := string(lower) + string(upper)
			for _, c := range string(r) {
				path = append(path, byte(c))
				backtrack(startIdx + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

func letterCasePermutation1(s string) []string {
	ans, path := []string{}, []byte{}
	n := len(s)
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if startIdx == n {
			str := string(path)
			ans = append(ans, str)
			return
		}
		backtrack(startIdx + 1)
		if unicode.IsDigit(rune(s[startIdx])) {
			path = append(path, s[startIdx])
			backtrack(startIdx + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

func main() {
	fmt.Println(letterCasePermutation1("3z4"))
}
