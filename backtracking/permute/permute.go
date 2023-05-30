package main

import "fmt"

func combine(n int, k int) [][]int {
	ans := [][]int{}
	path := []int{}
	var dfs func(n, k int, startidx int)
	dfs = func(n, k int, startidx int) {
		if len(path) == k {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}

		for i := startidx; i <= (n - (k - len(path)) + 1); i++ {
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, k, 1)
	return ans
}

func main() {
	fmt.Println(combine(4, 2))
}
