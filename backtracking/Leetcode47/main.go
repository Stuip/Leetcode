package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}
	used := make([]bool, n)
	sort.Ints(nums)
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			// 树层去重
			if i > 0 && used[i-1] && nums[i] == nums[i-1] {
				continue
			}
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				dfs(i + 1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
