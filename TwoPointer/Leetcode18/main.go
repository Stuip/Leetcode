package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					for l += 1; l < r && nums[l-1] == nums[l]; l++ {
					}
					for r -= 1; l < r && nums[r] == nums[r+1]; r-- {
					}
				} else if sum > target {
					for r -= 1; l < r && nums[r] == nums[r+1]; r-- {
					}
				} else {
					for l += 1; l < r && nums[l-1] == nums[l]; l++ {
					}
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
