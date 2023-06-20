package main

import "fmt"

func permute(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	path := make([]int, n)
	used := make([]bool, n)
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if startIdx == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for idx, on := range used {
			if !on {
				path[startIdx] = nums[idx]
				used[idx] = true
				backtrack(startIdx + 1)
				used[idx] = false
			}
		}
	}
	backtrack(0)
	return ans
}

func permute1(nums []int) (ans [][]int) {
	n := len(nums)
	path := []int{}
	used := make([]bool, n)
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				dfs(i + 1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute1(nums))

}
