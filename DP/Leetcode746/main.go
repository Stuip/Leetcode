package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	// f(i) = min(f(i-1)+cost[i-1], f(i-2)+cost[i-2])
	f0, f1 := 0, 0
	for i := 2; i <= len(cost); i++ {
		newF := min(f0+cost[i-2], f1+cost[i-1])
		f0, f1 = f1, newF
	}
	return f1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}
