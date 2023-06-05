package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	flag := false
	if x < 0 {
		flag = true
		x = -x
	}
	reverseNum := 0
	for x != 0 {
		if reverseNum < math.MinInt32/10 || reverseNum > math.MaxInt32/10 {
			return 0
		}
		reverseNum = reverseNum*10 + x%10
		x /= 10
	}
	if flag {
		reverseNum = -reverseNum
	}
	return reverseNum
}

func main() {
	fmt.Println(reverse(9646324351))
}
