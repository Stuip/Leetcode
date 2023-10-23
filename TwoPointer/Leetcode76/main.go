package main

import "fmt"

func minWindow(s string, t string) string {
	lenS, lenT := len(s), len(t)
	var sCount, tCount [52]int
	for i := 0; i < lenT; i++ {
		tCount[getIndex(t[i])] += 1
	}
	start, matchCount := 0, 0
	ans := ""
	for end := 0; end < lenS; end += 1 {
		indexEnd := getIndex(s[end])
		if tCount[indexEnd] > 0 && sCount[indexEnd] < tCount[indexEnd] {
			// 表示该字符在t中存在,并且s[left:right+1]之间还没有完全覆盖
			matchCount += 1
		}
		sCount[indexEnd] += 1
		for matchCount == lenT {
			if end-start+1 < len(ans) || len(ans) == 0 {
				ans = s[start : end+1]
			}
			indexStart := getIndex(s[start])
			if tCount[indexStart] > 0 && sCount[indexStart] <= tCount[indexStart] {
				// 如果s[left:right+1]的开始部分剔除,影响到匹配个数,则删减一个
				matchCount -= 1
			}
			sCount[indexStart] -= 1
			start++
		}

	}
	return ans
}

func getIndex(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}
	return int(c - 'A' + 26)
}

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow(s, t))
}
