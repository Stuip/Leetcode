package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix, suffix := make([]int, n), make([]int, n)
	prefix[0], suffix[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}
	for i := n - 2; i >= 0; i -= 1 {
		suffix[i] = suffix[i+1] * nums[i]
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			nums[i] = suffix[i+1]
			continue
		}
		if i == n-1 {
			nums[i] = prefix[i-1]
			continue
		}
		nums[i] = prefix[i-1] * suffix[i+1]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
