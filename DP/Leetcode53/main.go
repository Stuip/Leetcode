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

func maxSubArray2(nums []int) int {
	return getMax(nums, 0, len(nums)-1)
}

func getMax(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)/2
	leftMax := getMax(nums, left, mid)
	rightMax := getMax(nums, mid+1, right)
	crossMax := getCrossMax(nums, left, right)
	return max(max(leftMax, rightMax), crossMax)
}

func getCrossMax(nums []int, left, right int) int {
	mid := left + (right-left)/2
	leftSum := nums[mid]
	leftMax := leftSum
	for i := mid - 1; i >= left; i-- {
		leftSum += nums[i]
		leftMax = max(leftMax, leftSum)
	}
	rightSum := nums[mid+1]
	rightMax := rightSum
	for i := mid + 2; i <= right; i++ {
		rightSum += nums[i]
		rightMax = max(rightMax, rightSum)
	}
	return leftMax + rightMax
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
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
	fmt.Println(maxSubArray1(nums))
	fmt.Println(maxSubArray2(nums))
}
