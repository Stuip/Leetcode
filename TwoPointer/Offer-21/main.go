package main

import "fmt"

/**
Odd -> even
*/
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if !isOdd(nums[i]) && isOdd(nums[j]) {
			// 前面偶数， 后面奇数
			nums[i], nums[j] = nums[j], nums[i]
		} else if isOdd(nums[i]) && isOdd(nums[j]) {
			// 前面奇数， 后面奇数
			i += 1
		} else if !isOdd(nums[i]) && !isOdd(nums[j]) {
			// 前面偶数，后面偶数
			j -= 1
		} else if isOdd(nums[i]) && !isOdd(nums[j]) {
			// 前面奇数， 后面偶数
			i += 1
			j -= 1
		}
	}
	return nums
}

func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func exchange1(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && isOdd(nums[i]) {
			i += 1
		}
		for i < j && !isOdd(nums[j]) {
			j -= 1
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange1(nums))
}
