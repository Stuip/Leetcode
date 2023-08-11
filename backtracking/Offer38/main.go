package main

import "fmt"

func permutation(s string) []string {
	path := []byte{}
	n := len(s)
	used := make([]bool, n)
	m := make(map[string]bool)
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		if len(path) == n {
			m[string(path)] = true
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				path = append(path, s[i])
				dfs(i + 1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	ans := []string{}
	for k, _ := range m {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	s := "aab"
	fmt.Println(permutation(s))
}
