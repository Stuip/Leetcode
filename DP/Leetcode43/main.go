package main

import (
	"fmt"
)

/**
1.DP
	对于下标 i，下雨后水能到达的最大高度等于下标 i 两边的最大高度的最小值，
	下标 i 处能接的雨水量等于下标 i 处的水能到达的最大高度减去 height[i]。
*/

func trap_DP(height []int) int {
	n := len(height)
	preMax, sufMax := make([]int, n), make([]int, n)
	preMax[0], sufMax[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += min(sufMax[i], preMax[i]) - height[i]
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

/*
*
单调栈

	维护一个单调栈，单调栈存储的是下标，满足从栈底到栈顶的下标对应的数组 height 中的元素递减。
*/
func trap_stack(height []int) int {
	stack := []int{}
	ans := 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap_DP(height))
	fmt.Println(trap_stack(height))
}
