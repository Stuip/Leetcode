package main

import "fmt"

// BFS
func jump(nums []int) int {
	ans, n := 0, len(nums)
	used := make([]bool, n)
	queue := []int{0}
	used[0] = true
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			idx := queue[i]
			if idx == n-1 {
				return ans
			}
			for k := idx + 1; k <= idx+nums[idx] && k < n; k++ {
				if !used[k] {
					used[k] = true
					queue = append(queue, k)
				}
			}
		}
		queue = queue[sz:]
		ans += 1
	}
	return ans
}

// DP

func Jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j+nums[j] < i {
			j++
		}
		dp[i] = dp[j] + 1
	}
	return dp[n-1]
}

func main() {
	nums := []int{1, 2}
	fmt.Println(Jump(nums))
}
