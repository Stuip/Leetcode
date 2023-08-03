package main

import "sort"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 1
	for i >= 1 && nums[i-1] >= nums[i] {
		i -= 1
	}

	if i == 0 {
		sort.Ints(nums)
		return
	}
	i -= 1
	j := len(nums) - 1
	for nums[j] <= nums[i] {
		j -= 1
	}
	nums[i], nums[j] = nums[j], nums[i]
	arr := nums[i+1:]
	sort.Ints(arr)
	nums = append(nums[:i], arr...)
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
}
