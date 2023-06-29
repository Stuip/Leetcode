package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx < 0 {
			return 0
		}
		if cache[idx] != -1 {
			return cache[idx]
		}
		res := max(dfs(idx-1), dfs(idx-2)+nums[idx])
		cache[idx] = res
		return res
	}
	return dfs(n - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 时间复杂度：O(n) 空间复杂度： O(n)
func rob1(nums []int) int {
	n := len(nums)
	f := make([]int, n+2)
	for i, x := range nums {
		f[i+2] = max(f[i+1], f[i]+x)
	}
	return f[n+1]
}

// 时间复杂度：O(n) 空间复杂度： O(1)
func rob2(nums []int) int {
	f0, f1 := 0, 0
	for _, x := range nums {
		newF := max(f1, f0+x)
		f0, f1 = f1, newF
	}
	return f1
}

func rob3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = nums[0], 0
	for i := 1; i < n; i++ {
		// 第0列表示为打劫，第二列表示不打劫
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func main() {
	fmt.Println(rob2([]int{1, 2, 3, 1}))
}
