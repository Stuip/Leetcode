package main

import (
	"fmt"
)

// use DP[n][sum+1]
func canPartition(nums []int) bool {
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}

	// 如果该和为奇数，则无法成立
	if sum&1 == 1 {
		return false
	}

	sum /= 2
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, sum+1)
	}
	// 考虑第一个元素
	for i := sum; i >= 0; i-- {
		if i >= nums[0] {
			dp[0][i] = nums[0]
		}
	}
	// 第一层遍历，考虑第i个物品
	for i := 1; i < n; i++ {
		// 第二层遍历，考虑第当前背包还剩多少重量
		for j := sum; j >= 0; j-- {
			if j >= nums[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][sum] == sum
}

// use DP[2][sum+1]  使用滚动数组
func canPartition1(nums []int) bool {
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}

	// 如果该和为奇数，则无法成立
	if sum&1 == 1 {
		return false
	}

	sum /= 2
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, sum+1)
	}
	// 考虑第一个元素
	for i := sum; i >= 0; i-- {
		if i >= nums[0] {
			dp[0][i] = nums[0]
		}
	}
	// 第一层遍历，考虑第i个物品
	for i := 1; i < n; i++ {
		// 第二层遍历，考虑第当前背包还剩多少重量
		for j := sum; j >= 0; j-- {
			if j >= nums[i] {
				dp[i&1][j] = max(dp[(i-1)&1][j], dp[(i-1)&1][j-nums[i]]+nums[i])
			} else {
				dp[i&1][j] = dp[(i-1)&1][j]
			}
		}
	}
	return dp[(n-1)&1][sum] == sum
}

// O(n)
func canPartition2(nums []int) bool {
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}

	// 如果该和为奇数，则无法成立
	if sum&1 == 1 {
		return false
	}

	sum /= 2
	dp := make([]int, sum+1)
	// 考虑第一个元素
	for i := sum; i >= 0; i-- {
		if i >= nums[0] {
			dp[i] = nums[0]
		}
	}
	// 第一层遍历，考虑第i个物品
	for i := 1; i < n; i++ {
		// 第二层遍历，考虑第当前背包还剩多少重量
		for j := sum; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			}
		}
	}
	return dp[sum] == sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 5, 10, 6}
	fmt.Println(canPartition2(nums))
}
