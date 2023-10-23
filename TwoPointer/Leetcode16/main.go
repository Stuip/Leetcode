package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	ans := nums[0] + nums[1] + nums[2]
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > target {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for left, right := i+1, n-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if ans == target {
				return target
			} else if sum < target {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
