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

func subsetsWithDup1(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	sort.Ints(nums)
	var backtrack func(choosePre bool, startIdx int)
	backtrack = func(choosePre bool, startIdx int) {
		if startIdx == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		backtrack(false, startIdx+1)
		if !choosePre && startIdx > 0 && nums[startIdx-1] == nums[startIdx] {
			return
		}
		path = append(path, nums[startIdx])
		backtrack(true, startIdx+1)
		path = path[:len(path)-1]
	}
	backtrack(false, 0)
	return ans
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup1(nums))
}
