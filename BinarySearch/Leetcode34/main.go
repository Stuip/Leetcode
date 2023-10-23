package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				ans[0] = mid
				for mid < len(nums) && nums[mid] == target {
					mid += 1
				}
				ans[1] = mid - 1
				return ans
			}
			high = mid - 1
		}
	}
	return ans
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}
