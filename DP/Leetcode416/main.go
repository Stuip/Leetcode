package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

func canPartition1(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 0
	}
	for j := target; j >= 0; j-- {
		if j >= nums[0] {
			dp[0][j] = nums[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := target; j >= 1; j-- {
			if j >= nums[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target] == target
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 5, 10, 6}
	fmt.Println(canPartition1(nums))
}
