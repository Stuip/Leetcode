package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	path := []int{}
	var dfs func(target int, startIdx int)
	dfs = func(target int, startIdx int) {
		if len(path) == k && target == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		for i := startIdx; i < 10; i++ {
			if i > n {
				return
			}
			path = append(path, i)
			dfs(target-i, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, 1)
	return ans
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
