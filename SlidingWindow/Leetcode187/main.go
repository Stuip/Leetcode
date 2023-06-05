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
		if set[s[i:i+10]] > 1 {
			ans = append(ans, s[i:i+10])
		}
	}
	return ans
}

func main() {
	s := "AAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
