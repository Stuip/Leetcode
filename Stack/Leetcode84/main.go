package main

import "fmt"

func largestRectangleArea(heights []int) int {
	n := len(heights)
	l, r := make([]int, n), make([]int, n)
	for i:=0;i<n;i++{
		l[i] = -1
		r[i] = n
	}

	stack := []int{}

	for i:=0;i<n;i++{
		for len(stack) != 0 && heights[stack[len(stack)-1]] > heights[i] {
			r[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = stack[:0]

	for i:=n-1;i>=0;i--{
		for len(stack) != 0 && heights[stack[len(stack)-1]] > heights[i] {
			l[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	fmt.Println(l, r)

	ans := 0
	for i:=0;i<n;i++{
		h := heights[i]
		a, b := l[i], r[i]
		ans = max(ans, h * (b-a-1))
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
	nums := []int{2,1,5,6,2,3}
	fmt.Println(largestRectangleArea(nums))
}
