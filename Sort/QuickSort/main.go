package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, low, high int) {
	if high > low {
		pivot := partition(nums, low, high)
		QuickSort(nums, low, pivot-1)
		QuickSort(nums, pivot+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	for low < high {
		// high 指针值 >= pivot => high 会向右移
		for low < high && pivot <= nums[high] {
			high--
		}
		// 填补nums[low]
		nums[low] = nums[high]
		// low 指针的值 <= pivot => low 会左移
		for low < high && pivot >= nums[low] {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}
