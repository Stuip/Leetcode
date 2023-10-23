package main

import "fmt"

// 提意就是找到当前元素下一比它大的元素下标
func dailyTemperaturesFromFrontToBack(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperaturesFromBackToFront(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(stack) > 0 && t >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			idx := stack[len(stack)-1]
			ans[i] = idx - i
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperaturesFromFrontToBack(temperatures))
	fmt.Println(dailyTemperaturesFromBackToFront(temperatures))
}
