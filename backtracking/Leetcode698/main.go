package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum, n := 0, len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	// 如果这个和不能平均分配，则直接返回false
	if sum%k != 0 {
		return false
	}
	target := sum / k
	buckets := make([]int, k)
	sort.Ints(nums)
	for i, j := 0, n-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	var backtrack func(startIdx int) bool
	backtrack = func(startIdx int) bool {
		if startIdx == n {
			return true
		}
		for i := 0; i < k; i++ {
			if i > 0 && buckets[i] == buckets[i-1] {
				continue
			}
			if nums[startIdx]+buckets[i] > target {
				continue
			}
			buckets[i] += nums[startIdx]
			if backtrack(startIdx + 1) {
				return true
			}
			buckets[i] -= nums[startIdx]
		}
		return false
	}
	return backtrack(0)
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	fmt.Println(canPartitionKSubsets(nums, k))
}
