package main

import (
	"fmt"
)

func longestOnes(nums []int, k int) int {
	ans := 0
	start, count := 0, 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			count += 1
		}
		for count > k {
			if nums[start] == 0 {
				count -= 1
			}
			start += 1
		}
		ans = max(ans, end-start+1)
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
	nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	K := 3
	fmt.Println(longestOnes(nums, K))
}
