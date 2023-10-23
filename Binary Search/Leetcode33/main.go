package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}

	// 首先寻找到旋转点
	left, right := 0, n-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] >= nums[0] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if target >= nums[0] {
		left = 0
	} else {
		left = left + 1
		right = n - 1
	}
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if nums[right] == target {
		return right
	}
	return -1

}

func main() {
	nums := []int{1}
	target := 1
	fmt.Println(search(nums, target))
}
