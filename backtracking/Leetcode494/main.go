package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	var dfs func(start int, sum int)
	ans := 0
	dfs = func(start int, sum int) {
		if start > n {
			return
		}
		if start == n && sum == target {
			ans += 1
		}
		if start < n {
			dfs(start+1, sum+nums[start])
			dfs(start+1, sum-nums[start])
		}
	}
	dfs(0, 0)
	return ans
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}
