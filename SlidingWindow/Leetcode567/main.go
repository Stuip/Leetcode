package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	var v1, v2 [26]int
	l1, l2 := len(s1), len(s2)

	for i := 0; i < l1; i++ {
		v1[s1[i]-'a']++
	}
	for i := 0; i < l2; i++ {
		if i >= l1 {
			v2[s2[i-l1]-'a']--
		}
		v2[s2[i]-'a']++
		if v1 == v2 {
			return true
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
