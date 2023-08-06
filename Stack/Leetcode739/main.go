package main

import "fmt"

// 提意就是找到当前元素下一比它大的元素下标
func dailyTemperatures(temperatures []int) []int {
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

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
