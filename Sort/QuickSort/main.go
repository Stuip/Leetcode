package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

/**
快速排序： 基于比较，不稳定算法，时间复杂度为O(nlogn), 最坏情况下O(n ^ 2)
思想：
	分治思想，选主元，依次将剩余元素的小于主元放其左侧，大的放右侧;
	然后取主元的前半部分和后半部分进行同样处理，直至各子序列剩余一个元素结束，排序完成。
*/
func QuickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(nums, start, end)
	QuickSort(nums, start, mid-1)
	QuickSort(nums, mid+1, end)
}

func partition(nums []int, start, end int) int {
	randIdx := start + rand.Intn(end-start+1)
	nums[start], nums[randIdx] = nums[randIdx], nums[start]
	for start < end {
		for start < end && nums[start] < nums[end] {
			end -= 1
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start += 1
		}
		for start < end && nums[start] < nums[end] {
			start += 1
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
			end -= 1
		}
	}
	return start
}
