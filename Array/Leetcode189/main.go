package main

import "fmt"

func rotate(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, num := range nums {
		ans[(i+k)%n] = num
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Println(rotate(nums, k))
}
