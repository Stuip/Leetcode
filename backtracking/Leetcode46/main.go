package main

import "fmt"

func permute(nums []int) [][]int {
	ans, path := [][]int{}, []int{}
	n := len(nums)
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if len(path) == n {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		for i := startIdx; i < n; i++ {
			path = append(path, nums[i])
			backtrack(startIdx + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))

}
