package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	n, ans := len(nums), nums[k-1]-nums[0]
	for i := k; i < n; i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
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
	nums := []int{9, 4, 1, 7}
	k := 2
	fmt.Println(minimumDifference(nums, k))
}
