package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	m := make(map[int]bool, len(nums))
	for i, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
		if len(m) == k+1 {
			delete(m, nums[i-k])
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
