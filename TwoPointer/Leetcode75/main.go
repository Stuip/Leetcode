package main

import "fmt"

func sortColors(nums []int) []int {
	n := len(nums)
	i, j, idx := 0, n-1, 0
	for idx <= j {
		if nums[idx] == 0 {
			nums[idx], nums[i] = nums[i], nums[idx]
			i += 1
			idx += 1
		} else if nums[idx] == 2 {
			nums[idx], nums[j] = nums[j], nums[idx]
			idx += 1
		} else if nums[idx] == 1 {
			idx += 1
		}
	}
	return nums
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(sortColors(nums))
}
