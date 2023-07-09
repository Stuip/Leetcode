package main

import "fmt"

// DP
func trap(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = height[i]
		} else {
			left[i] = max(left[i-1], height[i])
		}
	}
	for i := n - 1; i >= 0; i -= 1 {
		if i == n-1 {
			right[i] = height[i]
		} else {
			right[i] = max(right[i+1], height[i])
		}
	}

	for i := 0; i < n; i++ {
		ans += min(left[i], right[i]) - height[i]
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

func trap1(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	preMax, sufMax := 0, 0
	for left <= right {
		preMax = max(height[left], preMax)
		sufMax = max(height[right], sufMax)
		if preMax < sufMax {
			ans += preMax - height[left]
			left += 1
		} else {
			ans += sufMax - height[right]
			right -= 1
		}
	}
	return ans
}

func main() {
	fmt.Println(trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
