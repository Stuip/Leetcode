package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	left, sum, ans := 0, 0, len(nums)+1
	for right, v := range nums {
		sum += v
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}
