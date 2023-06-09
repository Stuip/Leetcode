package main

import "fmt"

func removeElement(nums []int, val int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}

func removeElement1(nums []int, val int) int {
	j := len(nums) - 1
	for i := 0; i <= j; i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return j + 1
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
