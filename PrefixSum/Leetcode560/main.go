package main

import "fmt"

func subarraySum(nums []int, k int) int {
	n, counter := len(nums), 0
	pre := 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < n; i++ {
		pre += nums[i]
		if c, ok := m[pre-k]; ok {
			counter += c
		}
		m[pre] += 1
	}
	return counter
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}
