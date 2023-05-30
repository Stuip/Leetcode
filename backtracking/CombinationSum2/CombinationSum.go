package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	path := []int{}
	sort.Ints(candidates)
	var dfs func(target int, startidx int)
	dfs = func(target int, startidx int) {
		if target == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		for i := startidx; i < len(candidates); i++ {
			if candidates[i] > target {
				return
			}
			if i > startidx && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)
	return ans
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}
