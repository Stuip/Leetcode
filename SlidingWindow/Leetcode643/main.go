package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	ans := -math.MaxFloat64
	n, sum := 0, float64(0)
	for i, num := range nums {
		n += 1
		sum += float64(num)
		if n == k {
			if ave := sum / float64(k); ave > ans {
				ans = ave
			}
			sum -= float64(nums[i-k+1])
			n -= 1
		}
	}
	return ans
}

func main() {
	nums := []int{-1}
	n := 1
	fmt.Println(findMaxAverage(nums, n))
}
