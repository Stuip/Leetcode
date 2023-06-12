package main

import "fmt"

/**
哈系表
*/
func findRepeatNumber(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = true
		}
	}
	return -1
}

func findRepeatNumber1(nums []int) int {
	idx := 0
	for idx < len(nums) {
		if nums[idx] == idx {
			idx += 1
			continue
		}
		if nums[nums[idx]] == nums[idx] {
			return nums[idx]
		}
		nums[nums[idx]], nums[idx] = nums[idx], nums[nums[idx]]
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber1([]int{2, 3, 1, 0, 2, 5, 3}))
}
