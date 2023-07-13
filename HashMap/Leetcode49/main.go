package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	m := make(map[[128]byte]([]string))
	for _, str := range strs {
		var record [128]byte
		for i := 0; i < len(str); i++ {
			record[str[i]] += 1
		}
		m[record] = append(m[record], str)
	}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
