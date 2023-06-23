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

func main() {
	fmt.Println(rob2([]int{1, 2, 3, 1}))
}
