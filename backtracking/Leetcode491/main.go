package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		if len(path) > 1 {
			ans = append(ans, append([]int{}, path...))
		}
		used := make(map[int]bool, len(nums))
		for i := startIdx; i < len(nums); i++ {
			if used[nums[i]] { // 去重
				continue
			}
			if len(path) == 0 || path[len(path)-1] <= nums[i] {
				path = append(path, nums[i])
				used[nums[i]] = true
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func main() {
	nums := []int{4, 1, 7, 7}
	fmt.Println(findSubsequences(nums))
}
