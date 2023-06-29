package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	return max(helper(nums, 1, n), helper(nums, 0, n-1))
}

func helper(nums []int, low, high int) int {
	yes, no := nums[low], 0
	for i := low + 1; i < high; i++ {
		// 下一个状态中打劫和不打劫的最优解
		yes, no = no+nums[i], max(yes, no)
	}
	return max(yes, no)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1}))
}
