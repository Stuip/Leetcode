package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	ans := [][]int{}
	path := []int{}
	var dfs func(target int)
	dfs = func(target int) {
		if target == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		for i := 0; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(target - nums[i])
			path = path[:len(path)-1]
		}
	}
	dfs(target)
	return len(ans)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(combinationSum4(nums, 4))
}
