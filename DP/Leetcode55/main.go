package main

import "fmt"

func canJump(nums []int) bool {
	k := 0 // 覆盖范围
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		if k >= len(nums)-1 {
			return true
		}
		k = max(k, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 2, 1, 1, 0, 4}
	fmt.Println(canJump(nums))
}
