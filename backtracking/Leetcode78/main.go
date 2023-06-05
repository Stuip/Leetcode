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

func main() {
	nums := []int{0}
	fmt.Println(subsets(nums))
}
