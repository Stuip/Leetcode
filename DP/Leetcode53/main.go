package main

import (
	"fmt"
	"math"
)

// bruce method
func maxSubArray1(nums []int) int {
	n := len(nums)
	ans := math.MinInt
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			ans = max(ans, sum)
		}
	}
	return ans
}

func maxSubArray(nums []int) int {
	// dp[i]数组保存着， 以索引i为结尾的最大子序列和
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
