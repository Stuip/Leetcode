package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var targetP, record [128]int
	for _, c := range p {
		targetP[c-'a'] += 1
	}
	left, n := 0, len(s)
	m := len(p)
	ans := []int{}
	for right := 0; right < n; right++ {
		record[s[right]-'a'] += 1
		if right-left+1 == m {
			if targetP == record {
				ans = append(ans, left)
			}
			record[s[left]-'a'] -= 1
			left += 1
		}
	}
	return ans
}

func main() {
	s, p := "cbaebabacd", "abc"
	fmt.Println(findAnagrams(s, p))
}
