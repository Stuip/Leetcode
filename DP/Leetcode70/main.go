package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	f0, f1 := 0, 1
	for i := 0; i < n; i++ {
		newF := f0 + f1
		f0, f1 = f1, newF
	}
	return f1
}

func main() {
	fmt.Println(climbStairs(3))
}
