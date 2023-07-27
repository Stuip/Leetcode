package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := n + 1
	sums := make([]int, n+1)
	// 计算前缀和数组
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		s := sums[i]
		d := s - target
		left, right := 0, i
		for left < right {
			mid := (left + right + 1) >> 1
			if sums[mid] > d {
				right = mid - 1
			} else {
				left = mid
			}
		}
		if sums[right] <= d {
			ans = min(ans, i-right)
		}
	}

	if ans == n+1 {
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
