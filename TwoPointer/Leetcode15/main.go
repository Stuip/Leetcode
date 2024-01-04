package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
    n, ans := len(nums), [][]int{}
    sort.Ints(nums)
    for i:=0;i<n;i++{
        if nums[i] > 0 {
            break
        }

        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left, right := i+1, n-1
        for left < right {
            sum := nums[left] + nums[right] + nums[i]
            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[left], nums[right]})
                for left++;left<right&&nums[left]==nums[left-1];left++{}
                for right--;left<right&&nums[right]==nums[right+1];right--{}
            } else if sum > 0 {
                for left++;left<right&&nums[left]==nums[left-1];left++{}
            } else {
                for right--;left<right&&nums[right]==nums[right+1];right--{}
            }
        }
    }
    return ans
}

func main() {
	nums := []int{1,-1,-1,0}
	fmt.Println(threeSum(nums))
}