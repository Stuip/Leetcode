package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	ans := []string{}
	if len(s) < 10 {
		return ans
	}
	set := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		set[s[i:i+10]]++
	}
	for k, v := range set {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

func main() {
	s := "AAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
