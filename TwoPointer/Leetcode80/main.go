package main

import "fmt"

func removeDuplicates(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[ans] {
			ans += 1
			nums[ans] = nums[i]
		}
	}
	return ans
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
