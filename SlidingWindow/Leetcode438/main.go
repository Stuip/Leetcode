package main

import "fmt"

/*
*Note: 在golang中，基本类型可比较，数组类型需要保持相同长度和相同类型才比较。
 */
func findAnagrams(s string, p string) []int {
	ans := []int{}
	var vp [26]int
	var vs [26]int
	n, l := len(s), len(p)
	for _, c := range p {
		vp[c-'a']++
	}
	for i := 0; i < n; i++ {
		if i >= l {
			vs[s[i-l]-'a']--
		}
		vs[s[i]-'a']++
		if vs == vp {
			ans = append(ans, i-l+1)
		}
	}
	return ans
}

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}
