package main

import "fmt"

func maxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	for left < right {
		h := min(height[left], height[right])
		ans = max(ans, h*(right-left))
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(heights))
}
