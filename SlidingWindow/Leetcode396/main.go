package main

import (
	"fmt"
)

// 滑动窗口 + 前缀和
func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum := make([]int, n*2)
	sum[0] = nums[0]
	for i := 1; i < 2*n; i++ {
		sum[i] = sum[i-1] + nums[i%n]
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += nums[i] * i
	}
	for i, cur := n, ans; i < 2*n; i++ {
		//  每次向右移动之后，需要添加nums[(i-1)%n] * (n-1)
		cur += nums[i%n] * (n - 1)
		// 并且需要删除公共部分的值
		cur -= sum[i-1] - sum[i-n]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

func main() {
	nums := []int{4, 3, 2, 6}
	fmt.Println(maxRotateFunction(nums))
}
