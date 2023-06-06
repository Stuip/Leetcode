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
				break
			}
			path = append(path, candidates[i])
			dfs(target-candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)
	return ans
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
