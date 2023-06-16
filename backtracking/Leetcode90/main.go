package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	sort.Ints(nums)
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		ans = append(ans, append([]int{}, path...))
		if startIdx > len(nums) {
			return
		}
		for i := startIdx; i < len(nums); i++ {
			if i > startIdx && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
