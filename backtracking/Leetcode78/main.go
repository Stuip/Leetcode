package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	var backtracking func(startIdx int)
	backtracking = func(startIdx int) {
		copyPath := make([]int, len(path))
		copy(copyPath, path)
		ans = append(ans, copyPath)

		for i := startIdx; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}

func subsets1(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	var backtracking func(startIdx int)
	backtracking = func(startIdx int) {
		if startIdx == len(nums) {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		// 不选择nums[i], 直接跳过该数
		backtracking(startIdx + 1)
		path = append(path, nums[startIdx])
		backtracking(startIdx + 1)
		path = path[:len(path)-1]
	}
	backtracking(0)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
	fmt.Println(subsets1(nums))
}
