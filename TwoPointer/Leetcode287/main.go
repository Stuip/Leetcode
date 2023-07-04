package main

import "fmt"

/***
类似环形链表
*/
func findDuplicate(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
