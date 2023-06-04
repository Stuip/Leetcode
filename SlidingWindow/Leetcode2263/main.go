package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	numToStr := strconv.Itoa(num)
	start, ans := 0, 0
	for end := 0; end < len(numToStr); end++ {
		for end-start == k-1 {
			n, _ := strconv.Atoi(numToStr[start : end+1])
			if n != 0 && (num%n == 0) {
				ans += 1
			}
			start += 1
		}
	}
	return ans
}

func main() {
	num := 430043
	k := 2
	fmt.Println(divisorSubstrings(num, k))
}
