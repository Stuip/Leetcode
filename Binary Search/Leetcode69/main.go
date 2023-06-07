package main

import "fmt"

func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := low + (high-low)>>1
		num := mid * mid
		if num == x {
			return mid
		} else if num > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

func main() {
	fmt.Println(mySqrt(8))
}
