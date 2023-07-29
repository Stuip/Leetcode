package main

import "fmt"

func maxProduct(nums []int) int {
	ans := nums[0]
	preMax, preMin := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		a, b := preMax*nums[i], preMin*nums[i]
		preMax = max(nums[i], max(a, b))
		preMin = min(nums[i], min(a, b))
		ans = max(ans, preMax)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums))
}
